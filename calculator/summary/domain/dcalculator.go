package domain

import (
	mbill "github.com/tumpal2796/shopee/bill/model"
	"github.com/tumpal2796/shopee/calculator/summary/model"
)

type SummaryCalcInf interface {
	GetSummary(bills []mbill.DetailBill) model.Summary
}

type SummaryCalcImpl struct{}

func (summaryImpl *SummaryCalcImpl) GetSummary(bills []mbill.DetailBill) model.Summary {
	result := model.Summary{}

	for _, bill := range bills {
		result.PriceSubtotal += bill.Transaction.Price
		result.TaxSubtotal += bill.CalcDetail.Tax
		result.GrandTotal += bill.CalcDetail.Amount
	}

	return result
}
