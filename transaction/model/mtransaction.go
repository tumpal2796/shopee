package model

type Transaction struct {
	Name    string  `db:"name" json:"name"`
	TaxCode int8    `db:"tax_code" json:"tax_code"`
	Type    string  `db:"-" json:"type"`
	Price   float64 `db:"price" json:"price"`
}
