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

func (m *TransactionService) TotalBalance(marketName string) (float64, error) {
	result, err := m.dbService.MarketBalance(context.Background(), marketName)
	if err != nil {
		return 0, err
	}
	return result, nil

}

func (m *TransactionService) ListOperations(marketId int32) ([]postgres.ListMarketTransactionRow, error) {
	result, err := m.dbService.ListMarketTransaction(context.Background(), marketId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
