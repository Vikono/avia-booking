-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS companies
(
    id                   uuid primary key,
    name                 text                        not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE companies;
-- +goose StatementEnd
