-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refund_constraints
(
    id                   uuid                        not null,
    CONSTRAINT fk_id
        FOREIGN KEY (id)
            REFERENCES refund_policies (id),
    PRIMARY key (id, allowed_duration),
    cost                 int                         not null,
    allowed_duration     int                         not null,
    desctiption          text                        not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
