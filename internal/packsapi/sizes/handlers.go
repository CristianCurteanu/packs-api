package sizes

import (
	"github.com/CristiCurteanu/pack-api/internal/common/packager"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type packSizesHandlers struct {
	pkg packager.Packager
}

func NewPackSizesHandlers(pkg packager.Packager) packSizesHandlers {
	return packSizesHandlers{pkg}
}

func (ps packSizesHandlers) GetAllSizes(ctx *fiber.Ctx) error {
	response := lo.Map(ps.pkg.List(), func(ps packager.PackSize, _ int) PackSizeResponse {
		return PackSizeResponse{ps.Capacity}
	})

	return ctx.JSON(response)
}

func (ps packSizesHandlers) SetPackSizes(ctx *fiber.Ctx) error {
	packSizes := new(PackSizesRequest)

	err := ctx.BodyParser(packSizes)
	if err != nil {
		return err
	}

	err = ps.pkg.Set(lo.Map(packSizes.Data, func(ps PackSizeRequest, _ int) packager.PackSize {
		return packager.PackSize{
			Capacity: ps.Capacity,
		}
	}))

	return err
}
