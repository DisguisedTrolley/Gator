-- +goose Up

CREATE TABLE users (
	id UUID PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
