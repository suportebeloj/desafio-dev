package main

import (
	"github.com/suportebeloj/desafio-dev/internal/cmd/settings"
	"github.com/suportebeloj/desafio-dev/internal/core"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/frameworks/http/api"
	"github.com/suportebeloj/desafio-dev/internal/usecases"
	"log"
)

func main() {
	dbService := postgres.New(settings.DbConn.Conn)
	parser := core.NewTransactionParser()
	transactionService := usecases.NewTransactionService(dbService, parser)

	log.Println("transaction service running")

	httpService := api.NewHTTPApiService(transactionService, &api.HTTPServiceOptions{DbService: dbService, UseLogger: true})
	httpService.Run(":8000")

}
