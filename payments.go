package main

import (
	"errors"
	"maps"
	"slices"
	"time"
)

type Payment struct {
	Id        int
	Amount    float64
	Currency  string
	Status    string
	CreatedAt time.Time
}

type PaymentsHandler struct {
	Payments map[int]Payment
	Index    int
}

func NewPaymentsHandler() *PaymentsHandler {
	return &PaymentsHandler{
		Payments: make(map[int]Payment),
		Index:    0,
	}
}

func (p *PaymentsHandler) GetAll() []Payment {
	payments := slices.Collect(maps.Values(p.Payments))
	return payments
}

func (p *PaymentsHandler) Update(id int, payload Payment) (*Payment, error) {

	payment, ok := p.Payments[id]

	if !ok {
		return nil, errors.New("No such item")
	}

	payment.Amount = payload.Amount
	payment.Currency = payload.Currency
	payment.Status = payload.Status

	p.Payments[id] = payment

	return &payment, nil

}

func (p *PaymentsHandler) Delete(id int) error {
	_, ok := p.Payments[id]
	if !ok {
		return errors.New("No such item")
	}
	delete(p.Payments, id)
	return nil
}

func (p *PaymentsHandler) Add(payment *Payment) error {
	payment.CreatedAt = time.Now()
	payment.Id = p.Index
	p.Payments[p.Index] = *payment
	p.Index++

	return nil
}

func (p *PaymentsHandler) Get(id int) (*Payment, error) {
	payment, ok := p.Payments[id]

	if ok {
		return &payment, nil
	} else {
		return nil, errors.New("Item does not exist")
	}

}
