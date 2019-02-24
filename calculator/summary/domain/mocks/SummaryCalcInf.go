// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/tumpal2796/shopee/bill/model"
import summarymodel "github.com/tumpal2796/shopee/calculator/summary/model"

// SummaryCalcInf is an autogenerated mock type for the SummaryCalcInf type
type SummaryCalcInf struct {
	mock.Mock
}

// GetSummary provides a mock function with given fields: bills
func (_m *SummaryCalcInf) GetSummary(bills []model.DetailBill) summarymodel.Summary {
	ret := _m.Called(bills)

	var r0 summarymodel.Summary
	if rf, ok := ret.Get(0).(func([]model.DetailBill) summarymodel.Summary); ok {
		r0 = rf(bills)
	} else {
		r0 = ret.Get(0).(summarymodel.Summary)
	}

	return r0
}
