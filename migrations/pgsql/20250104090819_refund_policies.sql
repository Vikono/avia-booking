-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refund_policies
(
    id                   uuid primary key,
    company_id           uuid                        not null,
    type                 text                        not null,
    CONSTRAINT fk_company_id
        FOREIGN KEY (company_id)
            REFERENCES companies (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refund_policies;
-- +goose StatementEnd
