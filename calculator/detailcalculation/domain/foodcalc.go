package domain

import (
	"github.com/tumpal2796/shopee/calculator/detailcalculation/model"
	mtransaction "github.com/tumpal2796/shopee/transaction/model"
)

type Food struct{}

func (food *Food) CalculateDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc {
	var result model.DetailCalc
	result.Refundable = "YES"
	result.Tax = setPrecision(0.1 * transaction.Price)
	result.Amount = transaction.Price + result.Tax
	return result
}
