package packsapi

type PackResponse struct {
	Capacity int `json:"capacity"`
	Stored   int `json:"stored"`
}

type PacksResponse struct {
	Packs             []PackResponse `json:"packs"`
	PacksDescriptions []string       `json:"packsDescriptions"`
}
