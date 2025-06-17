package main

import (
	"github.com/CristiCurteanu/pack-api/internal/packsapi"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	apiV1 := app.Group("/api/v1")
	packsapi.RegisterHandlers(apiV1)

}
