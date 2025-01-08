-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cities
(
    id                   uuid primary key,
    name                 text                        not null,
    country_id           uuid                        not null,
    CONSTRAINT fk_country_id
        FOREIGN KEY (country_id)
            REFERENCES countries (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cities;
-- +goose StatementEnd
