package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/protocols"
	"io"
	"log"
	"strings"
)

type HTTPApiService struct {
	transactionService protocols.ITransactionService
	App                *fiber.App
	dbService          postgres.Querier
}

type HTTPServiceOptions struct {
	DbService postgres.Querier
}

func NewHTTPApiService(transactionService protocols.ITransactionService, option *HTTPServiceOptions) *HTTPApiService {
	s := &HTTPApiService{transactionService: transactionService}
	app := fiber.New()
	s.App = app
	s.routes()

	if option != nil {
		s.dbService = option.DbService
	}

	return s
}

func (H *HTTPApiService) routes() {
	group := H.App.Group("/api/v1/")
	group.Add("POST", "new", H.CreateTransaction)
}

func (H *HTTPApiService) Run(addrs string) error {
	H.App.Use(logger.New())
	log.Fatalln(H.App.Listen(addrs))
	return nil
}

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

func (H *HTTPApiService) ListMarkets(c *fiber.Ctx) error {
	return nil

}

func (H *HTTPApiService) MarketDetail(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (H HTTPApiService) MarketBalance(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
