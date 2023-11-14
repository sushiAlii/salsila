CREATE TYPE gender_enum AS ENUM ('Male', 'Female');
CREATE TYPE family_role_enum AS ENUM ('Father', 'Mother', 'Child');

CREATE TABLE IF NOT EXISTS roles (
	id 	serial PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	description VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS social_networks (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	base_url VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS families (
	id SERIAL PRIMARY KEY,
	family_name VARCHAR(100) NOT NULL,
	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS persons (
	uid UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	first_name VARCHAR(100) NOT NULL,
	middle_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	gender gender_enum NOT NULL,
	birthday DATE NOT NULL,
	created_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS users (
	uid UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	role_id INTEGER REFERENCES roles(id) NOT NULL,
	persons_uid UUID REFERENCES persons(uid) UNIQUE,
	email VARCHAR(100) UNIQUE,
	password BYTEA NOT NULL,
	deleted_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ,
	created_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
	id SERIAL PRIMARY KEY,
	user_uid UUID REFERENCES users(uid) NOT NULL,
	token TEXT NOT NULL,
	expires_at TIMESTAMPTZ NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_networks (
	id SERIAL PRIMARY KEY,
	user_uid UUID REFERENCES users(uid) NOT NULL,
	social_networks_id INTEGER REFERENCES social_networks(id) NOT NULL,
	user_name VARCHAR(100),
	user_url VARCHAR(200) NOT NULL,
	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS persons_family (
	id SERIAL PRIMARY KEY,
	family_id INTEGER REFERENCES families(id) NOT NULL,
	person_uid UUID REFERENCES persons(uid) NOT NULL,
	family_role family_role_enum NOT NULL,
	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ
);
