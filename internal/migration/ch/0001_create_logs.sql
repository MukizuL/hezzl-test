-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS logs
(
    Id UInt32,
    ProjectId UInt32,
    Name String,
    Description String,
    Priority UInt32,
    Removed UInt8,
    EventTime DateTime
)
    ENGINE = MergeTree()
        ORDER BY (Id, ProjectId, Name);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS logs;