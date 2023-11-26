DROP TABLE IF EXISTS users CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE users
(
    id         UUID PRIMARY KEY,
    first_name VARCHAR(32)                   NOT NULL,
    last_name  VARCHAR(32)                   NOT NULL,
    email      VARCHAR(64) UNIQUE            NOT NULL,
    password   VARCHAR(64)                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE      NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    last_logged_at TIMESTAMP WITH TIME ZONE
);