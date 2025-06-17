package packager

import (
	"slices"
	"sort"
)

type PackSize struct {
	Capacity int
}

type PackSizes []PackSize

func (p PackSizes) Len() int {
	return len(p)
}

func (p PackSizes) Less(i int, j int) bool {
	return p[i].Capacity < p[j].Capacity
}

func (p PackSizes) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

type Pack struct {
	Capacity    int
	StoredItems int
}

type Packager interface {
	SetPackSizes(s PackSizes)
	PackItems(items int) []Pack
}

type packager struct {
	packSizes PackSizes
}

func NewPackager(s PackSizes) Packager {
	p := &packager{}

	p.SetPackSizes(s)

	return p
}

// SetPackSizes implements Packager.
func (p *packager) SetPackSizes(s PackSizes) {
	if s.Len() < 1 {
		panic("should be at least one pack sizes definition")
	}

	sCopy := make(PackSizes, s.Len())
	copy(sCopy, s)

	if !sort.IsSorted(sCopy) {
		sort.Sort(sCopy)
	}

	slices.Reverse(sCopy)
	p.packSizes = sCopy

}

// PackItems implements Packager.
func (p *packager) PackItems(items int) []Pack {
	panic("unimplemented")
}
