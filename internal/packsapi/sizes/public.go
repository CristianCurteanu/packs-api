package sizes

type PackSizeRequest struct {
	Capacity int `json:"capacity"`
}

type PackSizesRequest struct {
	Data []PackSizeRequest `json:"data"`
}

type PackSizeResponse struct {
	Capacity int `json:"capacity"`
}

type PackSizesResponse struct {
	Data []PackSizeResponse `json:"data"`
}
