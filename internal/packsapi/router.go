package packsapi

import "github.com/gofiber/fiber/v2"

func RegisterHandlers(api fiber.Router) {
	packsHandlers := newPacksHandlers()

	api.Get("/packs", packsHandlers.GetPacks)
}
