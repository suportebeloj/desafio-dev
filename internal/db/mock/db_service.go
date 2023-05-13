package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
)

type DbService struct {
	mock.Mock
}

func NewDbService() *DbService {
	return &DbService{}
}

func (m *DbService) CreateTransaction(ctx context.Context, arg postgres.CreateTransactionParams) (postgres.CreateTransactionRow, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(postgres.CreateTransactionRow), args.Error(1)
}

func (m *DbService) GetTransaction(ctx context.Context, id int32) (postgres.GetTransactionRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(postgres.GetTransactionRow), args.Error(1)
}

func (m *DbService) ListMarketTransaction(ctx context.Context, market string) ([]postgres.ListMarketTransactionRow, error) {
	args := m.Called(ctx, market)
	return args.Get(0).([]postgres.ListMarketTransactionRow), args.Error(0)
}

func (m *DbService) ListMarkets(ctx context.Context) ([]string, error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}
