package entity

import "github.com/google/uuid"

type TransactionType int

const (
	TransactionCredit TransactionType = iota + 1
	TransactionDebit
)

func (t TransactionType) String() string {
	switch t {
	case TransactionCredit:
		return "Credito"
	case TransactionDebit:
		return "Debito"
	}
	return "Desconhecido"
}

type Transaction struct {
	ID    uuid.UUID
	Type  TransactionType
	Value float64
}
