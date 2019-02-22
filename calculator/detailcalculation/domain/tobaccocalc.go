package domain

import (
	"math"

	"github.com/sophee/calculator/detailcalculation/model"
	mtransaction "github.com/sophee/transaction/model"
)

type Tobacco struct{}

func (tobacco *Tobacco) CalculateDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc {
	var result model.DetailCalc
	result.Refundable = "No"
	result.Tax = math.Round(10 + (0.02 * transaction.Price))
	result.Amount = transaction.Price + result.Tax
	return result
}
