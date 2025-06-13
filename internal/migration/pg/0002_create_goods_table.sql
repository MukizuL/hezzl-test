-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goods (
    id SERIAL,
    project_id INT,
    name TEXT NOT NULL,
    description TEXT,
    priority INT DEFAULT 0 NOT NULL,
    removed BOOL DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    PRIMARY KEY (id, project_id),
    CONSTRAINT fk_projects FOREIGN KEY (project_id) REFERENCES projects(id)
    );

CREATE INDEX goods_name_hash_idx ON goods USING HASH(name);

CREATE OR REPLACE FUNCTION set_priority()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.priority = (SELECT COALESCE(MAX(priority), 0) + 1 FROM goods);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_set_priority
    BEFORE INSERT ON goods
    FOR EACH ROW EXECUTE FUNCTION set_priority();
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS goods;