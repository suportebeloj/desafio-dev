package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "github.com/suportebeloj/desafio-dev/docs"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/protocols"
	"io"
	"log"
	"net/url"
	"strings"
)

type HTTPApiService struct {
	transactionService protocols.ITransactionService
	App                *fiber.App
	dbService          postgres.Querier
}

type HTTPServiceOptions struct {
	DbService postgres.Querier
	UseLogger bool
}

func NewHTTPApiService(transactionService protocols.ITransactionService, option *HTTPServiceOptions) *HTTPApiService {
	s := &HTTPApiService{transactionService: transactionService}
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "http://127.0.0.1:5173, http://localhost:5173,"}))
	if option.UseLogger == true {
		app.Use(logger.New())
	}
	s.App = app
	s.routes()

	if option.DbService != nil {
		s.dbService = option.DbService
	}

	return s
}

func (H *HTTPApiService) routes() {
	group := H.App.Group("/api/v1/")
	group.Add("POST", "new", H.CreateTransaction)
	group.Add("GET", "markets", H.ListMarkets)
	group.Add("GET", "detail/:market", H.MarketDetail)
	H.Swagger()
}

func (H *HTTPApiService) Run(addrs string) error {
	H.App.Use(logger.New())
	log.Fatalln(H.App.Listen(addrs))
	return nil
}

// CreateTransaction godoc
// @Summary      Create a Transaction
// @Description  upload a CNAB file to save the transactions on database
// @Accept       multipart/form-data
// @Produce      text/plain; charset=utf-8
// @Param        transactions   formData      file  true  "CNAB file"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /new [post]
func (H *HTTPApiService) CreateTransaction(c *fiber.Ctx) error {
	file, err := c.FormFile("transactions")
	if err != nil {
		log.Println("form", err)
		return err
	}

	fileContent, err := file.Open()
	if err != nil {
		log.Println("read", err)
		return err
	}
	defer fileContent.Close()

	content, err := io.ReadAll(fileContent)
	if err != nil {
		return err
	}

	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		_, _ = H.transactionService.NewTransaction(row)
	}

	return c.SendStatus(200)
}

// ListMarkets godoc
// @Summary      List Registered stores
// @Description  return a json contains a list of string representing the registered stores
// @Accept       text/plain; charset=utf-8
// @Produce      application/json
// @Success      200
// @Failure      500
// @Router       /markets [get]
func (H *HTTPApiService) ListMarkets(c *fiber.Ctx) error {
	result, err := H.dbService.ListMarkets(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(result)
}

// MarketDetail godoc
// @Summary      Market Detail
// @Description  return a json contains information about the store and your transactions
// @Accept       text/plain; charset=utf-8
// @Produce      application/json
// @Param        market   path      string  true  "market name"
// @Success      200
// @Failure      404
// @Failure      500
// @Router       /detail/{market} [get]
func (H *HTTPApiService) MarketDetail(c *fiber.Ctx) error {
	market := c.Params("market")
	market, err := url.QueryUnescape(market)
	if err != nil {
		return err
	}

	result, err := H.transactionService.ListOperations(market)
	if err != nil {
		return err
	}

	balance, err := H.dbService.MarketBalance(context.Background(), market)
	if err != nil {
		return err
	}

	if balance < 0 {
		balance = balance * -1
	}

	if len(result) > 0 {
		info := struct {
			MarketName string                              `json:"market_name"`
			Owner      string                              `json:"owner"`
			Balance    float64                             `json:"balance"`
			Operations []postgres.ListMarketTransactionRow `json:"operations"`
		}{
			MarketName: result[0].Market,
			Owner:      result[0].Owner,
			Balance:    balance,
			Operations: result,
		}

		return c.JSON(info)
	}

	return c.Status(fiber.StatusNotFound).SendString("No transactions found for this store")

}

func (H *HTTPApiService) Swagger() {
	H.App.Get("/swagger/*", swagger.HandlerDefault)
	H.App.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "http://localhost:8000/docs.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))
}
