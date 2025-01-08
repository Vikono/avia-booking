-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS airports
(
    id                   uuid primary key,
    airport_code         text                      not null,
    city_id              uuid                        not null,
    CONSTRAINT fk_city_id
        FOREIGN KEY (city_id)
            REFERENCES cities (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE airports;
-- +goose StatementEnd
