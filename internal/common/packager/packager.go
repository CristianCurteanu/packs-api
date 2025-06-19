package packager

import (
	"errors"
	"slices"
	"sort"

	slicesInternal "github.com/CristiCurteanu/pack-api/internal/common/slices"
)

type Packager interface {
	Set(s PackSizes) error
	List() PackSizes
	PackItems(items int) []Pack
}

type packager struct {
	packSizes PackSizes
}

func (p *packager) List() PackSizes {
	return p.packSizes
}

func NewPackager(s PackSizes) Packager {
	p := &packager{}

	err := p.Set(s)
	if err != nil {
		panic(err)
	}

	return p
}

func (p *packager) Set(s PackSizes) error {
	if s.Len() < 1 {
		return errors.New("should be at least one pack sizes definition")
	}

	sCopy := make(PackSizes, s.Len())
	copy(sCopy, s)

	if !sort.IsSorted(sCopy) {
		sort.Sort(sCopy)
	}

	slices.Reverse(sCopy)
	p.packSizes = sCopy

	return nil
}

func (p *packager) PackItems(items int) []Pack {
	var result []Pack
	remaining := items

	for remaining != 0 {
		var prev int
		sizes := slicesInternal.Filter(p.packSizes, func(s PackSize, i int) bool {
			prev = i - 1
			return s.Capacity <= remaining
		})

		var maxSize PackSize
		if len(sizes) != 0 {
			maxSize = sizes[0]
		} else {
			maxSize = p.packSizes[len(p.packSizes)-1]
		}

		pack := Pack{Capacity: maxSize.Capacity}

		if maxSize.Capacity < remaining {
			pack.StoredItems = maxSize.Capacity
			remaining -= maxSize.Capacity
		} else {
			// Squeeze the last package if there is capacity to last packages
			if len(result) > 0 && result[len(result)-1].Capacity+remaining < p.packSizes[prev].Capacity {
				pack.Capacity = p.packSizes[prev].Capacity
				pack.StoredItems = result[len(result)-1].Capacity + remaining
				result = result[:len(result)-1]
			} else {
				pack.StoredItems = remaining
			}
			remaining = 0
		}

		result = append(result, pack)
	}

	return result
}
