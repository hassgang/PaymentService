package repository

import (
	"errors"

	. "payment_api/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PaymentPg struct {
	db *sqlx.DB
}

func NewPaymentPg(db *sqlx.DB) PaymentRepo {
	return &PaymentPg{db: db}
}

type PaymentRepo interface {
	Create(p *Payment) (*Payment, error)
	Get(id int64) (*Payment, error)
	GetAll() ([]Payment, error)
	Update(id int64, p *PaymentUpdate) (*Payment, error)
	Delete(id int64) error
}

func (repo *PaymentPg) Create(p *Payment) (*Payment, error) {
	err := repo.db.QueryRow(`insert into payments values (default,$1,$2,$3,$4,$5) returning id`, p.Amount, p.Currency, p.Status, p.CreatedAt, p.IdempotencyKey).Scan(&p.Id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return nil, ErrAlreadyExists
		}
		return nil, err
	}
	return p, nil
}

func (repo *PaymentPg) Update(id int64, p *PaymentUpdate) (*Payment, error) {
	var payment Payment
	err := repo.db.QueryRowx("update payments set status=$1 where id=$2 returning *", p.Status, id).StructScan(&payment)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (repo *PaymentPg) Get(id int64) (*Payment, error) {
	var p Payment
	err := repo.db.QueryRowx("select * from payments where id=$1", id).StructScan(&p)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (repo *PaymentPg) GetAll() ([]Payment, error) {
	var payments []Payment
	err := repo.db.Select(&payments, "select * from payments")

	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (repo *PaymentPg) Delete(id int64) error {
	res, err := repo.db.Exec("delete from payments where id=$1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return nil
	}

	return nil

}
