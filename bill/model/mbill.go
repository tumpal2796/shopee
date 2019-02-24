package model

import (
	mdetailbill "github.com/tumpal2796/shopee/calculator/detailcalculation/model"
	msummary "github.com/tumpal2796/shopee/calculator/summary/model"
	mtransaction "github.com/tumpal2796/shopee/transaction/model"
)

type Bill struct {
	DetailBills []DetailBill     `json:"detail_bills"`
	Summary     msummary.Summary `json:"summary"`
}

type DetailBill struct {
	Transaction mtransaction.Transaction `json:"transaction"`
	CalcDetail  mdetailbill.DetailCalc   `json:"calc_detail"`
}
