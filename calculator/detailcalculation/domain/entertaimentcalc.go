package domain

import (
	"math"

	"github.com/sophee/calculator/detailcalculation/model"
	mtransaction "github.com/sophee/transaction/model"
)

type Entertaiment struct{}

func (entertaiment *Entertaiment) CalculateDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc {
	var result model.DetailCalc
	result.Refundable = "NO"
	if transaction.Price > 0 && transaction.Price < 100 {
		result.Tax = 0
	} else if transaction.Price >= 100 {
		result.Tax = math.Round(0.01 * (transaction.Price - 100))
	}
	result.Amount = transaction.Price + result.Tax

	return result
}
