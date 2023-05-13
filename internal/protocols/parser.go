package protocols

import "github.com/suportebeloj/desafio-dev/internal/db/postgres"

type ITransactionParser interface {
	Parser(transactionString string) (*postgres.CreateTransactionParams, error)
}
