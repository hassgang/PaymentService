-- +goose Up
create table if not exists Payments(
    id SERIAL PRIMARY KEY,
    amount  DECIMAL(12,2) NOT NULL,
    currency text not null,
    status TEXT NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMPTZ NOT NULL,
    idempotency_key UUID NOT NULL UNIQUE
);

-- +goose Down
drop table if exists Payments;
