package packs

import (
	"net/http"
	"strconv"

	"github.com/CristiCurteanu/pack-api/internal/common/packager"
	"github.com/gofiber/fiber/v2"
)

type packsHandlers struct {
	pkg packager.Packager
}

func NewPacksHandlers(pkg packager.Packager) packsHandlers {
	return packsHandlers{pkg}
}

func (ph packsHandlers) GetPacks(ctx *fiber.Ctx) error {

	queryParams := ctx.Queries()

	itemString, paramFound := queryParams["items"]
	if !paramFound {
		return fiber.NewError(http.StatusBadRequest, "{\"error\": \"no items query params given\"}")
	}

	items, err := strconv.Atoi(itemString)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(
		mapResponse(ph.pkg.PackItems(items)),
	)
}
