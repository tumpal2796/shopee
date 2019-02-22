package domain

import (
	"github.com/tumpal2796/sophee/calculator/detailcalculation/model"
	mtransaction "github.com/tumpal2796/sophee/transaction/model"
)

type Tobacco struct{}

func (tobacco *Tobacco) CalculateDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc {
	var result model.DetailCalc
	result.Refundable = "NO"
	result.Tax = setPrecision(10 + (0.02 * transaction.Price))
	result.Amount = transaction.Price + result.Tax
	return result
}
