package protocols

import "github.com/gofiber/fiber/v2"

type IHTTPApiService interface {
	Run(addrs string) error
	CreateTransaction(c *fiber.Ctx) error
	ListMarkets(c *fiber.Ctx) error
	MarketDetail(c *fiber.Ctx) error
	MarketBalance(c *fiber.Ctx) error
}
