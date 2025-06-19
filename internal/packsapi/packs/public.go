package packs

import (
	"fmt"

	"github.com/CristiCurteanu/pack-api/internal/common/packager"
)

type PackResponse struct {
	Capacity int `json:"capacity"`
	Stored   int `json:"stored"`
}

type PacksResponse struct {
	Packs             []PackResponse `json:"packs"`
	PacksDescriptions []string       `json:"packsDescriptions"`
}

func mapResponse(packs []packager.Pack) PacksResponse {
	var resp PacksResponse
	descriptions := make(map[int]int)

	for _, pack := range packs {
		resp.Packs = append(resp.Packs, PackResponse{
			Capacity: pack.Capacity,
			Stored:   pack.StoredItems,
		})
		descriptions[pack.Capacity] += 1
	}

	for c, reps := range descriptions {
		resp.PacksDescriptions = append(resp.PacksDescriptions, fmt.Sprintf("%d X %d", reps, c))
	}

	return resp
}
