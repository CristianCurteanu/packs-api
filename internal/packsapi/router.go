package packsapi

import (
	"github.com/CristiCurteanu/pack-api/internal/common/packager"
	"github.com/CristiCurteanu/pack-api/internal/packsapi/packs"
	"github.com/CristiCurteanu/pack-api/internal/packsapi/sizes"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandlers(api fiber.Router, pkgr packager.Packager) {
	packsHandlers := packs.NewPacksHandlers(pkgr)
	api.Get("/packs", packsHandlers.GetPacks)

	packSizesHandlers := sizes.NewPackSizesHandlers(pkgr)
	api.Get("/packs/sizes", packSizesHandlers.GetAllSizes)
	api.Put("/packs/sizes", packSizesHandlers.SetPackSizes)
}
