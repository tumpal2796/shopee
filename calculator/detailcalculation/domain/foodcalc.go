package domain

import (
	"math"

	"github.com/sophee/calculator/detailcalculation/model"
	mtransaction "github.com/sophee/transaction/model"
)

type Food struct{}

func (food *Food) CalculateDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc {
	var result model.DetailCalc
	result.Refundable = "Yes"
	result.Tax = math.Round(0.1 * transaction.Price)
	result.Amount = transaction.Price + result.Tax
	return result
}
