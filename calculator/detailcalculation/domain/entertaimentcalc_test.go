package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tumpal2796/shopee/calculator/detailcalculation/model"
	mtransaction "github.com/tumpal2796/shopee/transaction/model"
)

func TestCalculateDetailCalculationEntertainment(t *testing.T) {
	assert := assert.New(t)
	type TestCase struct {
		transaction    mtransaction.Transaction
		expectedOutput model.DetailCalc
	}

	testCases := make(map[string]TestCase)
	testCases["Get Detail Calucation of Entertainment With No Tax"] = TestCase{
		transaction: mtransaction.Transaction{
			Price: 50,
		},
		expectedOutput: model.DetailCalc{
			Refundable: "NO",
			Tax:        0,
			Amount:     50,
		},
	}

	testCases["Get Detail Calucation of Entertainment With No Tax2"] = TestCase{
		transaction: mtransaction.Transaction{
			Price: 100,
		},
		expectedOutput: model.DetailCalc{
			Refundable: "NO",
			Tax:        0,
			Amount:     100,
		},
	}

	testCases["Get Detail Calucation of Entertainment With Tax"] = TestCase{
		transaction: mtransaction.Transaction{
			Price: 1000,
		},
		expectedOutput: model.DetailCalc{
			Refundable: "NO",
			Tax:        9,
			Amount:     1009,
		},
	}

	entertainment := &Entertaiment{}
	for key, val := range testCases {
		result := entertainment.CalculateDetailCalculation(val.transaction)
		assert.Equal(val.expectedOutput, result, "[TestCalculateDetailCalculationEntertainment] Failed on Test %s", key)
	}

	t.Logf("[TestCalculateDetailCalculationEntertainment] Success")
	return
}
