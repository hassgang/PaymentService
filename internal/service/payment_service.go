package service

import (
	. "payment_api/internal/model"
	. "payment_api/internal/repository"
	"time"
)

type PaymentsService struct {
	repo PaymentRepo
}

type Service interface {
	Create(p *PaymentCreate) (*Payment, error)
	Get(id int64) (*Payment, error)
	GetAll() ([]Payment, error)
	Update(id int64, p *PaymentUpdate) (*Payment, error)
	Delete(id int64) error
}

func NewService(repo PaymentRepo) Service {
	return &PaymentsService{repo: repo}
}

func (s *PaymentsService) GetAll() ([]Payment, error) {
	payments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (s *PaymentsService) Update(id int64, pu *PaymentUpdate) (*Payment, error) {

	p, err := s.repo.Update(id, pu)

	if err != nil {
		return nil, err
	}
	return p, nil

}

func (s *PaymentsService) Delete(id int64) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentsService) Create(pc *PaymentCreate) (*Payment, error) {

	p := Payment{
		CreatedAt:      time.Now(),
		Amount:         pc.Amount,
		Currency:       pc.Currency,
		IdempotencyKey: pc.IdempotencyKey,
		Status:         "Pending",
	}
	res, err := s.repo.Create(&p)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PaymentsService) Get(id int64) (*Payment, error) {

	p, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}
	return p, nil

}
