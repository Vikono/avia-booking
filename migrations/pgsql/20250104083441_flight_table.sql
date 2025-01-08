-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flights
(
    id                   uuid primary key,
    origin_id            uuid                        not null,
    CONSTRAINT fk_origin_id
        FOREIGN KEY (origin_id)
            REFERENCES airports (id),
    destination_id       uuid                        not null,
    CONSTRAINT fk_destination_id
        FOREIGN KEY (destination_id)
            REFERENCES airports (id),
    departure_date       date                        not null,
    departure_time       time                        not null,
    arrival_date         date                        not null,
    arrival_time         time                        not null,
    company_id           uuid                        not null,
    CONSTRAINT fk_company_id
        FOREIGN KEY (company_id)
            REFERENCES companies (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE flights;
-- +goose StatementEnd
