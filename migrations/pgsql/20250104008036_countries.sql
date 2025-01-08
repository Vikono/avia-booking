-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS countries
(
    id                   uuid primary key,
    name                 text                        not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE countries;
-- +goose StatementEnd
