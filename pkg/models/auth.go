package models

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	accessSecret = os.Getenv("ACCESS_SECRET")
	refreshSecret = os.Getenv("REFESH_SECRET")
)

type TokenDetails struct {
	AccessToken 	string
	RefreshToken	string
	ATExpires		int64
	RTExpires		int64
}

type RefreshToken struct {
	ID			uint		`gorm:"primaryKey"`
	UserUID		string		`gorm:"type:uuid;not null"`
	Token		string		`gorm:"not null"`
	ExpiresAt	time.Time	`gorm:"type:timestamptz;not null"`
	CreatedAt	time.Time	`gorm:"type:timestamptz;not null"`
}

type AuthService interface {
	LoginUser(string, string) (*User, error)
	LogoutUser(string) error
	RegisterUser(*User) error
	CreateToken(string) (*TokenDetails, error)
	SaveAuth(string, *TokenDetails) error
	Refresh(string) (*TokenDetails, error)
}

type authService struct {
	DB *gorm.DB
	userService UserService
}

func NewAuthService(db *gorm.DB, userService UserService) AuthService {
	return &authService{DB: db, userService: userService}
}

func (as *authService) LoginUser(email string, password string) (*User, error) {
	user, err := as.userService.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrPasswordIncorrect
	}

	return user, nil
}

func (as *authService) LogoutUser(refreshToken string) error {
	userUid, err := getUserUIDFromToken(refreshToken)
	if err != nil {
		return err
	}

	tx := as.DB.Begin()

	if err := tx.Where("user_uid = ? AND token = ?", userUid, refreshToken).Delete(&RefreshToken{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (as *authService) RegisterUser(user *User) error {
	return as.userService.CreateUser(user)
}

func (as *authService) CreateToken(userUID string) (*TokenDetails, error) {
	token := &TokenDetails{}

	token.ATExpires = time.Now().Add(time.Minute * 15).Unix()
	token.RTExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_uid"] = userUID
	atClaims["exp"] = token.ATExpires

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = accessToken.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["user_uid"] = userUID
	rtClaims["exp"] = token.RTExpires

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = refreshToken.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (as *authService) SaveAuth(userUid string, td *TokenDetails) error {
	token := &RefreshToken{
		Token:		td.RefreshToken,
		UserUID: 	userUid,
		ExpiresAt: 	time.Unix(td.RTExpires, 0),
	}

	return as.DB.Create(token).Error;
}

func (as *authService) Refresh(refreshToken string) (*TokenDetails, error) {
	tx := as.DB.Begin()

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	userUid, ok := claims["user_uid"].(string)
	if !ok {
		return nil, fmt.Errorf("Invalid claim: %s", userUid)
	}

	if _, err := verifyToken(as.DB, userUid, refreshToken); err != nil {
		return nil, err
	}

	newToken, err := as.CreateToken(userUid)
	if err != nil {
		return nil, err
	}

	if err := as.SaveAuth(userUid, newToken); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	} 

	return newToken, nil
}

func verifyToken(DB *gorm.DB, userUid string, tokenString string) (*RefreshToken, error) {
	var token RefreshToken

	if err := DB.Where("user_uid = ? AND token = ?", userUid, tokenString).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Token not found")
		}
		return nil, err
	}

	if token.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("Token has expired")
	}

	if err := DB.Where("user_uid = ? AND token = ?", userUid, tokenString).Delete(&RefreshToken{}).Error; err != nil {
		return nil, fmt.Errorf("Unable to refresh token: %v", err)
	}

	return &token, nil
}

func getUserUIDFromToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(refreshSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
		userUid, ok := claims["user_uid"].(string)
		if !ok {
			return "", fmt.Errorf("user_uid not found in token")
		}

		return userUid, nil
	} else {
		return "", fmt.Errorf("Invalid token")
	}
}