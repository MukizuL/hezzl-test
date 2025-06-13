-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
    );

INSERT INTO projects (name) VALUES ('Первая запись');
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS projects;