CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    bio VARCHAR(255) NOT NULL,
    beatdrops INT NOT NULL DEFAULT 0,
    friends INT NOT NULL DEFAULT 0,
    settings JSON NOT NULL,
    photo TEXT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS beats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    userid UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT now(),
    location VARCHAR(255) NOT NULL,
    song VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    longitude FLOAT NOT NULL,
    latitude FLOAT NOT NULL,
    image TEXT NOT NULL,
    comments INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    userid UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    beatid UUID NOT NULL REFERENCES beats(id) ON DELETE CASCADE,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT now(),
    comment VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS friends (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    alpha UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    beta UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT now(),
    status INT NOT NULL DEFAULT 0,
    sender UUID NOT NULL,
    CONSTRAINT user_pair_unique UNIQUE (alpha, beta),
    CONSTRAINT user_pair_check CHECK (alpha > beta)
);