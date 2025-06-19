package packager

var DefaultPackSizes = PackSizes{
	{Capacity: 5000},
	{Capacity: 500},
	{Capacity: 2000},
	{Capacity: 250},
	{Capacity: 1000},
}

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
