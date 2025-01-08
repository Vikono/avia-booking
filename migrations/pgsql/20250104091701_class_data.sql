-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS class_data
(
    flight_id            uuid                        not null,
    CONSTRAINT fk_flight_id
        FOREIGN KEY (flight_id)
            REFERENCES flights (id),
    class_type           text                        not null,
    primary key (flight_id, class_type),
    cost                 int                         not null,
    luggage              int                         not null,
    buggage              int                         not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE class_data;
-- +goose StatementEnd
