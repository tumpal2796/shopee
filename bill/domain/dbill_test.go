package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tumpal2796/shopee/bill/model"
	dcalcdetail "github.com/tumpal2796/shopee/calculator/detailcalculation/domain"
	calcdetailmocks "github.com/tumpal2796/shopee/calculator/detailcalculation/domain/mocks"
	mcalcdetail "github.com/tumpal2796/shopee/calculator/detailcalculation/model"
	dsummary "github.com/tumpal2796/shopee/calculator/summary/domain"
	summarymocks "github.com/tumpal2796/shopee/calculator/summary/domain/mocks"
	msummary "github.com/tumpal2796/shopee/calculator/summary/model"
	mtransaction "github.com/tumpal2796/shopee/transaction/model"
	"github.com/tumpal2796/shopee/transaction/resource"
	tmocks "github.com/tumpal2796/shopee/transaction/resource/mocks"
)

func TestNewBill(t *testing.T) {
	assert := assert.New(t)
	transactionmock := new(tmocks.TransactionInf)

	type TestCase struct {
		transactionRes resource.TransactionInf
		expOutput      *BillImpl
	}

	testCase := TestCase{
		transactionRes: transactionmock,
		expOutput: &BillImpl{
			TransactionRes:   transactionmock,
			CalculatorDetail: &dcalcdetail.CalculatorImpl{},
			Summary:          &dsummary.SummaryCalcImpl{},
		},
	}

	result := NewBill(testCase.transactionRes)
	assert.Equal(testCase.expOutput, result, "[TestNewBill] Failed")

	t.Logf("[TestNewBill] Success")
	return
}

func TestGetMyBill(t *testing.T) {
	assert := assert.New(t)
	transacactionMock := new(tmocks.TransactionInf)
	calculationDetailMock := new(calcdetailmocks.CalculatorInf)
	summaryMock := new(summarymocks.SummaryCalcInf)
	billImpl := &BillImpl{
		TransactionRes:   transacactionMock,
		CalculatorDetail: calculationDetailMock,
		Summary:          summaryMock,
	}

	expectedOutput := model.Bill{
		DetailBills: []model.DetailBill{
			model.DetailBill{
				Transaction: mtransaction.Transaction{
					Name:    "oreo",
					TaxCode: 1,
					Type:    "food",
					Price:   1000,
				},
				CalcDetail: mcalcdetail.DetailCalc{
					Refundable: "YES",
					Tax:        100,
					Amount:     1100,
				},
			},
		},
		Summary: msummary.Summary{
			PriceSubtotal: 3000,
			TaxSubtotal:   200,
			GrandTotal:    3200,
		},
	}

	transacactionMock.On("GetAllTransaction", mock.Anything).Return([]mtransaction.Transaction{
		mtransaction.Transaction{
			Name:    "oreo",
			TaxCode: 1,
			Type:    "food",
			Price:   1000,
		},
	}, nil)

	calculationDetailMock.On("GetDetailCalculation", mock.Anything).Return(mcalcdetail.DetailCalc{
		Refundable: "YES",
		Tax:        100,
		Amount:     1100,
	})

	summaryMock.On("GetSummary", mock.Anything).Return(msummary.Summary{
		PriceSubtotal: 3000,
		TaxSubtotal:   200,
		GrandTotal:    3200,
	})

	result, err := billImpl.GetMyBill()
	assert.NoError(err, "[TestGetMyBill] Failed")
	assert.Equal(expectedOutput, result, "[TestGetMyBill] Failed")
	transacactionMock.AssertExpectations(t)
	calculationDetailMock.AssertExpectations(t)
	summaryMock.AssertExpectations(t)
	t.Logf("[TestGetMyBill] Success")
	return
}

func TestErrorGetMyBill(t *testing.T) {
	assert := assert.New(t)
	transacactionMock := new(tmocks.TransactionInf)
	calculationDetailMock := new(calcdetailmocks.CalculatorInf)
	summaryMock := new(summarymocks.SummaryCalcInf)
	billImpl := &BillImpl{
		TransactionRes:   transacactionMock,
		CalculatorDetail: calculationDetailMock,
		Summary:          summaryMock,
	}

	var expectedOutput model.Bill

	transacactionMock.On("GetAllTransaction", mock.Anything).Return(nil, errors.New("Failed to Get Transaction"))
	result, err := billImpl.GetMyBill()
	assert.Error(err, "[TestGetMyBill] Failed")
	assert.Equal(expectedOutput, result, "[TestGetMyBill] Failed")
	transacactionMock.AssertExpectations(t)

	t.Logf("[TestGetMyBill] Success")
	return
}
