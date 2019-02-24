package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mbill "github.com/tumpal2796/shopee/bill/model"
	mdetailbill "github.com/tumpal2796/shopee/calculator/detailcalculation/model"
	"github.com/tumpal2796/shopee/calculator/summary/model"
	mtransaction "github.com/tumpal2796/shopee/transaction/model"
)

func TestGetSummary(t *testing.T) {
	assert := assert.New(t)

	type TestCase struct {
		bills          []mbill.DetailBill
		expectedResult model.Summary
	}

	testCases := make(map[string]TestCase)
	testCases["Get  Summary Data"] = TestCase{
		bills: []mbill.DetailBill{
			mbill.DetailBill{
				Transaction: mtransaction.Transaction{
					Price: 1000,
				},
				CalcDetail: mdetailbill.DetailCalc{
					Tax:    100,
					Amount: 1100,
				},
			},
			mbill.DetailBill{
				Transaction: mtransaction.Transaction{
					Price: 2000,
				},
				CalcDetail: mdetailbill.DetailCalc{
					Tax:    200,
					Amount: 2200,
				},
			},
		},
		expectedResult: model.Summary{
			PriceSubtotal: 3000,
			TaxSubtotal:   300,
			GrandTotal:    3300,
		},
	}

	summaryCaclImpl := &SummaryCalcImpl{}
	for key, val := range testCases {
		result := summaryCaclImpl.GetSummary(val.bills)
		assert.Equal(val.expectedResult, result, "[TestGetSummary] Failed on Test %s", key)
	}

	t.Logf("[TestGetSummary] Success")
	return
}
