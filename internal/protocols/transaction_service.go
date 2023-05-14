package protocols

import "github.com/suportebeloj/desafio-dev/internal/db/postgres"

type ITransactionService interface {
	NewTransaction(transaction string) (postgres.CreateTransactionRow, error)
	TotalBalance(marketName string) (float64, error)
	ListOperations(marketId int32) ([]postgres.ListMarketTransactionRow, error)
}
