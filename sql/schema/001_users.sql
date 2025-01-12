-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    name text UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;