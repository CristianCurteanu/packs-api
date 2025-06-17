package packsapi

import "github.com/gofiber/fiber/v2"

type packsHandlers struct {
}

func newPacksHandlers() packsHandlers {
	return packsHandlers{}
}

func (ph packsHandlers) GetPacks(*fiber.Ctx) error {

	return nil
}
