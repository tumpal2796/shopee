package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tumpal2796/shopee/calculator/detailcalculation/model"
	mtransaction "github.com/tumpal2796/shopee/transaction/model"
)

func TestCalculateDetailCalculation(t *testing.T) {
	assert := assert.New(t)
	type TestCase struct {
		transaction    mtransaction.Transaction
		expectedOutput model.DetailCalc
	}

	testCases := make(map[string]TestCase)
	testCases["Food"] = TestCase{
		transaction: mtransaction.Transaction{
			TaxCode: 1,
			Price:   1000,
		},
		expectedOutput: model.DetailCalc{
			Refundable: "YES",
			Tax:        100,
			Amount:     1100,
		},
	}

	testCases["Tobacco"] = TestCase{
		transaction: mtransaction.Transaction{
			TaxCode: 2,
			Price:   1000,
		},
		expectedOutput: model.DetailCalc{
			Refundable: "NO",
			Tax:        30,
			Amount:     1030,
		},
	}

	testCases["Entertainment"] = TestCase{
		transaction: mtransaction.Transaction{
			TaxCode: 3,
			Price:   50,
		},
		expectedOutput: model.DetailCalc{
			Refundable: "NO",
			Tax:        0,
			Amount:     50,
		},
	}

	entertainment := &CalculatorImpl{}
	for key, val := range testCases {
		result := entertainment.GetDetailCalculation(val.transaction)
		assert.Equal(val.expectedOutput, result, "[TestCalculateDetailCalculation] Failed on Test %s", key)
	}

	t.Logf("[TestCalculateDetailCalculation] Success")
	return
}

func TestSetPrecision(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		input     float64
		expOutput float64
	}{
		{0.5, 0.50},
		{122.234, 122.23},
	}

	for key, val := range testCases {
		result := setPrecision(val.input)
		assert.Equal(val.expOutput, result, "[TestSetPrecision] Failed on Test %d", key)
	}

	t.Logf("[TestSetPrecision] Success")
	return
}
