package domain

import (
	"context"
	"time"

	"github.com/sophee/bill/model"
	dcalcdetail "github.com/sophee/calculator/detailcalculation/domain"
	dsummary "github.com/sophee/calculator/summary/domain"
	"github.com/sophee/transaction/resource"
	tresource "github.com/sophee/transaction/resource"
)

type BillInf interface {
	GetMyBill() (model.Bill, error)
}

type BillImpl struct {
	TransactionRes   tresource.TransactionInf
	CalculatorDetail dcalcdetail.CalculatorInf
	Summary          dsummary.SummaryCalcInf
}

func NewBill(TransactionRes resource.TransactionInf) BillInf {
	return &BillImpl{
		TransactionRes:   TransactionRes,
		CalculatorDetail: &dcalcdetail.CalculatorImpl{},
		Summary:          &dsummary.SummaryCalcImpl{},
	}
}

func (billImpl *BillImpl) GetMyBill() (model.Bill, error) {
	var result model.Bill
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	transactionData, err := billImpl.TransactionRes.GetAllTransaction(ctx)
	if err != nil {
		return result, err
	}

	for _, transaction := range transactionData {
		calculationDetail := billImpl.CalculatorDetail.GetDetailCalculation(transaction)
		result.DetailBills = append(result.DetailBills, model.DetailBill{
			Transaction: transaction,
			CalcDetail:  calculationDetail,
		})
	}

	summary := billImpl.Summary.GetSummary(result.DetailBills)
	result.Summary = summary
	return result, err
}
