package usecase

import "github.com/eminetto/immudb/entity"

type Bank interface {
	AddTransaction(t entity.Transaction) error
	GetBalance() (float64, error)
}

type BankStorage interface {
	SaveTransaction(t entity.Transaction) error
}
