package cmc

type Convert struct {
	Data *Result `json:"data"`
}

type Result struct {
	Quotes map[string]*struct {
		Price float64 `json:"price"`
	} `json:"quote"`
}
