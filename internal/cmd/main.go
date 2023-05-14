package main

import (
	"github.com/suportebeloj/desafio-dev/internal/cmd/settings"
	"github.com/suportebeloj/desafio-dev/internal/core"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/frameworks/http/api"
	"github.com/suportebeloj/desafio-dev/internal/usecases"
	"log"
	"os"
)

// @title           Swagger Desafio Dev
// @version         1.0
// @description     Test server for dev challenge.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Carlos
// @contact.url    https://github.com/suportebeloj
// @contact.email  carlos.e.alves2@proton.me

// @host      localhost:8000
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	dbService := postgres.New(settings.DbConn.Conn)
	parser := core.NewTransactionParser()
	transactionService := usecases.NewTransactionService(dbService, parser)

	log.Println("transaction service running")

	httpService := api.NewHTTPApiService(transactionService, &api.HTTPServiceOptions{DbService: dbService, UseLogger: true})
	httpService.Run(os.Getenv("PORT"))

}
