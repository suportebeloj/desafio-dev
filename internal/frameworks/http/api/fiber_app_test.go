package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockService "github.com/suportebeloj/desafio-dev/internal/db/mock"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/frameworks/http/api"
	"github.com/suportebeloj/desafio-dev/internal/usecases"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type StubParser struct {
}

func (s *StubParser) Parser(transactionString string) (*postgres.CreateTransactionParams, error) {
	return &postgres.CreateTransactionParams{}, nil
}

func TestHTTPApiSerice_CreateTransaction(t *testing.T) {

	fileName := "file.txt"
	fileField := "transactions"
	fileContent := `2201903010000010700845152540738723****9987123333MARCOS PEREIRAMERCADO DA AVENIDA\n
2201903010000050200845152540738473****1231231233MARCOS PEREIRAMERCADO DA AVENIDA\n
3201903010000060200232702980566777****1313172712JOSÉ COSTA    MERCEARIA 3 IRMÃOS
`

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part, err := writer.CreateFormFile(fileField, fileName)
	assert.NoError(t, err)

	_, err = io.Copy(part, strings.NewReader(fileContent))

	writer.Close()

	req := httptest.NewRequest("POST", "/api/v1/new", payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	dbService := mockService.NewDbService()
	dbService.On("CreateTransaction", mock.Anything, mock.Anything).Return(postgres.CreateTransactionRow{}, nil)
	parser := &StubParser{}
	transactionService := usecases.NewTransactionService(dbService, parser)
	httpService := api.NewHTTPApiService(transactionService, &api.HTTPServiceOptions{DbService: dbService})

	resp, _ := httpService.App.Test(req)

	assert.Equal(t, resp.StatusCode, fiber.StatusOK)

	dbService.AssertExpectations(t)

}

func TestHTTPApiService_CreateTransaction_ReturnError_WhenNotSendFile_WithRequest(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.Close()

	req := httptest.NewRequest("POST", "/api/v1/new", payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	dbService := mockService.NewDbService()
	dbService.On("CreateTransaction", mock.Anything, mock.Anything).Return(postgres.CreateTransactionRow{}, nil)
	parser := &StubParser{}
	transactionService := usecases.NewTransactionService(dbService, parser)
	httpService := api.NewHTTPApiService(transactionService, &api.HTTPServiceOptions{DbService: dbService})

	resp, err := httpService.App.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, resp.StatusCode, fiber.StatusInternalServerError)

}

func TestHTTPApiService_ListMarkets_ReturnAValidList_ContainsValidMarkets(t *testing.T) {
	expected := []string{
		"market 1",
		"market 2",
		"market 3",
		"market 4",
	}
	dbService := mockService.NewDbService()
	calledFunc := dbService.On("ListMarkets", mock.Anything).Return(expected, nil)
	parser := &StubParser{}
	transactionService := usecases.NewTransactionService(dbService, parser)
	httpService := api.NewHTTPApiService(transactionService, &api.HTTPServiceOptions{DbService: dbService})

	req := httptest.NewRequest("GET", "/api/v1/markets", nil)

	resp, err := httpService.App.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, fiber.StatusOK)

	defer resp.Body.Close()

	byteBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var body []string
	err = json.Unmarshal(byteBody, &body)
	assert.NoError(t, err)

	assert.Equal(t, expected, body)

	dbService.AssertExpectations(t)

	calledFunc.Unset()

	dbService.On("ListMarkets", mock.Anything).Return([]string{}, errors.New("database error"))

	req = httptest.NewRequest("GET", "/api/v1/markets", nil)

	resp, err = httpService.App.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, fiber.StatusInternalServerError)

}

func TestHTTPApiService_MarketDetail_ReceiveAValidStore_AndReturnADetailedStruct(t *testing.T) {

	type responseBody struct {
		MarketName string                              `json:"market_name"`
		Owner      string                              `json:"owner"`
		Balance    float64                             `json:"balance"`
		Operations []postgres.ListMarketTransactionRow `json:"operations"`
	}

	operations := []postgres.ListMarketTransactionRow{
		{
			ID:     1,
			Type:   "1",
			Date:   time.Now(),
			Value:  222,
			Cpf:    "00000000000",
			Card:   "000000000000000",
			Time:   time.Now(),
			Owner:  "Test Owner",
			Market: "Test store",
		},
		{
			ID:     2,
			Type:   "3",
			Date:   time.Now(),
			Value:  333,
			Cpf:    "00200000000",
			Card:   "000f20000000000",
			Time:   time.Now(),
			Owner:  "Test Owner",
			Market: "Test store",
		},
	}
	expected := responseBody{
		MarketName: "Test store",
		Owner:      "Test Owner",
		Balance:    123456,
		Operations: operations,
	}

	req := httptest.NewRequest("GET", "/api/v1/detail/Test%20store", nil)

	dbService := mockService.NewDbService()
	calledFunc := dbService.On("ListMarketTransaction", mock.Anything, mock.Anything).Return(operations, nil)
	dbService.On("MarketBalance", mock.Anything, mock.Anything).Return(-123456.0, nil)
	parser := &StubParser{}
	transactionService := usecases.NewTransactionService(dbService, parser)
	httpService := api.NewHTTPApiService(transactionService, &api.HTTPServiceOptions{DbService: dbService})

	resp, err := httpService.App.Test(req)
	assert.NoError(t, err)

	defer resp.Body.Close()

	byteBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	body := responseBody{}
	err = json.Unmarshal(byteBody, &body)
	assert.NoError(t, err)

	assert.Equal(t, expected.MarketName, body.MarketName)
	assert.Equal(t, expected.Owner, body.Owner)
	assert.Equal(t, expected.Balance, body.Balance)
	assert.Equal(t, expected.Operations[1].Market, body.Operations[1].Market)
	assert.Equal(t, expected.Operations[1].Value, body.Operations[1].Value)

	dbService.AssertExpectations(t)

	calledFunc.Unset()

	dbService.On("ListMarketTransaction", mock.Anything, mock.Anything).Return([]postgres.ListMarketTransactionRow{}, nil)

	req = httptest.NewRequest("GET", "/api/v1/detail/Test%20store", nil)
	resp, err = httpService.App.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

}
