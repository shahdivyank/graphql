CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    bio VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS beats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    userid UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    timestamp INT NOT NULL,
    location VARCHAR(255) NOT NULL,
    song VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    longitude FLOAT NOT NULL,
    latitude FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    userid UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    beatid UUID NOT NULL REFERENCES beats(id) ON DELETE CASCADE,
    timestamp INT NOT NULL,
    comment VARCHAR(255) NOT NULL
);