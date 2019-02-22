package model

type DetailCalc struct {
	Refundable string  `json:"refundable"`
	Tax        float64 `json:"tax"`
	Amount     float64 `json:"amount"`
}
