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
