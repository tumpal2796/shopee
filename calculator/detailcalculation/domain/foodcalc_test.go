package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tumpal2796/shopee/calculator/detailcalculation/model"
	mtransaction "github.com/tumpal2796/shopee/transaction/model"
)

func TestCalculateDetailCalculationFood(t *testing.T) {
	assert := assert.New(t)
	type TestCase struct {
		transaction    mtransaction.Transaction
		expectedOutput model.DetailCalc
	}

	testCases := make(map[string]TestCase)
	testCases["Get Detail Calucation of Food"] = TestCase{
		transaction: mtransaction.Transaction{
			Price: 1000,
		},
		expectedOutput: model.DetailCalc{
			Refundable: "YES",
			Tax:        100,
			Amount:     1100,
		},
	}

	entertainment := &Food{}
	for key, val := range testCases {
		result := entertainment.CalculateDetailCalculation(val.transaction)
		assert.Equal(val.expectedOutput, result, "[TestCalculateDetailCalculationFood] Failed on Test %s", key)
	}

	t.Logf("[TestCalculateDetailCalculationFood] Success")
	return
}
