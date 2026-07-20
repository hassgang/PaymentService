package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrAlreadyExists = errors.New("payment already exists")

type Payment struct {
	Id             int64
	Amount         float64
	Currency       string
	Status         string
	CreatedAt      time.Time `db:"created_at"`
	IdempotencyKey uuid.UUID `db:"idempotency_key"`
}

type PaymentCreate struct {
	Amount         float64
	Currency       string
	IdempotencyKey uuid.UUID
}

type PaymentUpdate struct {
	Status string
}
