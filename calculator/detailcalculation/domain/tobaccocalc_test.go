package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tumpal2796/sophee/calculator/detailcalculation/model"
	mtransaction "github.com/tumpal2796/sophee/transaction/model"
)

func TestCalculateDetailCalculationTobacco(t *testing.T) {
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
			Refundable: "NO",
			Tax:        30,
			Amount:     1030,
		},
	}

	entertainment := &Tobacco{}
	for key, val := range testCases {
		result := entertainment.CalculateDetailCalculation(val.transaction)
		assert.Equal(val.expectedOutput, result, "[TestCalculateDetailCalculationTobacco] Failed on Test %s", key)
	}

	t.Logf("[TestCalculateDetailCalculationTobacco] Success")
	return
}
