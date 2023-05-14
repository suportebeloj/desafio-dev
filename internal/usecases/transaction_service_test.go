package usecases_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/suportebeloj/desafio-dev/internal/core"
	testMock "github.com/suportebeloj/desafio-dev/internal/db/mock"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/usecases"
	"testing"
	"time"
)

func TestTransactionService_NewTransaction_ParseTransactionString_AndPersists(t *testing.T) {

	dbService := testMock.NewDbService()

	expected := postgres.CreateTransactionRow{
		ID:     1,
		Type:   "2",
		Date:   time.Date(2019, 3, 1, 0, 0, 0, 0, time.Local),
		Value:  502,
		Cpf:    "84515254073",
		Card:   "8473****1231",
		Time:   time.Date(0, 0, 0, 23, 12, 33, 0, time.Local),
		Owner:  "MARCOS PEREIRA",
		Market: "MERCADO DA AVENIDA",
	}

	dbService.On("CreateTransaction", mock.Anything, mock.Anything).Return(expected, nil)

	parser := core.NewTransactionParser()

	instance := usecases.NewTransactionService(dbService, parser)
	assert.NotNil(t, instance)
	transactionString := "2201903010000050200845152540738473****1231231233MARCOS PEREIRAMERCADO DA AVENIDA"

	result, err := instance.NewTransaction(transactionString)
	assert.NoError(t, err)
	assert.Equal(t, result, expected)

	dbService.AssertExpectations(t)

}

func TestTransactionService_TotalBalance(t *testing.T) {
	dbService := testMock.NewDbService()
	expected := 10.0
	calledFunc := dbService.On("MarketBalance", mock.Anything, mock.Anything).Return(expected, nil)
	parser := core.NewTransactionParser()

	instance := usecases.NewTransactionService(dbService, parser)

	result, err := instance.TotalBalance("test market")
	assert.NoError(t, err)

	assert.Equal(t, result, expected)

	dbService.AssertExpectations(t)

	calledFunc.Unset()

	dbService.On("MarketBalance", mock.Anything, mock.Anything).Return(0.0, errors.New("invalid market"))

	result, err = instance.TotalBalance("invalid test market")
	assert.Error(t, err, "invalid market")

	assert.Equal(t, result, 0.0)

}

func TestTransactionService_ListOperations(t *testing.T) {
	dbService := testMock.NewDbService()
	expected := []postgres.ListMarketTransactionRow{
		{ID: 1, Type: "1", Date: time.Now(), Value: 123, Cpf: "123445412", Card: "123324423", Time: time.Now(), Owner: "Test owner", Market: "Test Market"},
		{ID: 1, Type: "2", Date: time.Now(), Value: 321, Cpf: "423423434", Card: "534523f34", Time: time.Now(), Owner: "Test owner", Market: "Test Market"},
		{ID: 1, Type: "3", Date: time.Now(), Value: 223, Cpf: "213423423", Card: "34345t455", Time: time.Now(), Owner: "Test owner", Market: "Test Market"},
	}

	calledFunc := dbService.On("ListMarketTransaction", mock.Anything, mock.Anything).Return(expected, nil)

	parser := core.NewTransactionParser()

	instance := usecases.NewTransactionService(dbService, parser)
	results, err := instance.ListOperations(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, results)

	dbService.AssertExpectations(t)

	calledFunc.Unset()
	dbService.On("ListMarketTransaction", mock.Anything, mock.Anything).Return([]postgres.ListMarketTransactionRow{}, errors.New("invalid market name"))

	results, err = instance.ListOperations(2)
	assert.Error(t, err)
	assert.Nil(t, results)
}
