package domain

import (
	"math"

	"github.com/tumpal2796/sophee/calculator/detailcalculation/model"
	mtransaction "github.com/tumpal2796/sophee/transaction/model"
)

type CalculatorInf interface {
	GetDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc
}

type CalculateInf interface {
	CalculateDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc
}

type CalculatorImpl struct{}

func (calculatorImpl *CalculatorImpl) GetDetailCalculation(transaction mtransaction.Transaction) model.DetailCalc {
	var Tipe CalculateInf

	switch transaction.TaxCode {
	case 1:
		Tipe = &Food{}
	case 2:
		Tipe = &Tobacco{}
	case 3:
		Tipe = &Entertaiment{}
	}

	return Tipe.CalculateDetailCalculation(transaction)
}

func setPrecision(number float64) float64 {
	return math.Floor(number*100) / 100
}
