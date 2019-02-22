package model

type Summary struct {
	PriceSubtotal float64 `json:"price_sub_total"`
	TaxSubtotal   float64 `json:"tax_sub_total"`
	GrandTotal    float64 `json:"grand_total"`
}
