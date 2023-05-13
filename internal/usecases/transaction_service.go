package usecases

import (
	"context"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/protocols"
)

type TransactionService struct {
	dbService postgres.Querier
	parser    protocols.ITransactionParser
}

func NewTransactionService(dbService postgres.Querier, parser protocols.ITransactionParser) *TransactionService {
	return &TransactionService{dbService: dbService, parser: parser}
}

func (m *TransactionService) NewTransaction(transaction string) (postgres.CreateTransactionRow, error) {
	parsed, err := m.parser.Parser(transaction)
	if err != nil {
		return postgres.CreateTransactionRow{}, err
	}

	stroredTransaction, err := m.dbService.CreateTransaction(context.Background(), *parsed)
	if err != nil {
		return postgres.CreateTransactionRow{}, err
	}

	return stroredTransaction, nil

}

func (m *TransactionService) TotalBalance(marketName string) float64 {
	//TODO implement me
	panic("implement me")
}

func (m *TransactionService) ListOperations(id int) []postgres.Transaction {
	//TODO implement me
	panic("implement me")
}
