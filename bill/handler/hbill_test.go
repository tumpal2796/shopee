package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/julienschmidt/httprouter"
	"github.com/tumpal2796/sophee/bill/domain/mocks"
	"github.com/tumpal2796/sophee/bill/model"
	mdetailbill "github.com/tumpal2796/sophee/calculator/detailcalculation/model"
	msummary "github.com/tumpal2796/sophee/calculator/summary/model"
	mtransaction "github.com/tumpal2796/sophee/transaction/model"
)

func TestNewBill(t *testing.T) {
	assert := assert.New(t)
	domainBillMock := new(mocks.BillInf)
	expResult := &BillImpl{
		DomainBill: domainBillMock,
	}
	result := NewBill(domainBillMock)
	assert.Equal(expResult, result, "[TestNewBill] Failed")

	t.Logf("[TestNewBill] Success")
	return
}

func TestGetMyBill(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost:8008/getmybill", nil)
	p := httprouter.Params{}

	domainBillMock := new(mocks.BillInf)
	handlerBill := &BillImpl{
		DomainBill: domainBillMock,
	}

	domainBillMock.On("GetMyBill", mock.Anything, mock.Anything, mock.Anything).Return(model.Bill{
		DetailBills: []model.DetailBill{
			model.DetailBill{
				Transaction: mtransaction.Transaction{
					Name:    "oreo",
					TaxCode: 1,
					Type:    "food",
					Price:   1000,
				},
				CalcDetail: mdetailbill.DetailCalc{
					Refundable: "YES",
					Tax:        100,
					Amount:     1100,
				},
			},
		},
		Summary: msummary.Summary{
			PriceSubtotal: 1000,
			TaxSubtotal:   100,
			GrandTotal:    1100,
		},
	}, nil)

	handlerBill.GetMyBill(w, r, p)
	var result Response
	bodyBytes, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "[TestGetMyBill] Failed")
	err = json.Unmarshal(bodyBytes, &result)
	assert.NoError(err, "[TestGetMyBill] Failed")

	assert.Equal(result.StatusCode, 200, "[TestAddMyBill] Failed")
	t.Logf("[TestGetMyBill] Success")
	return
}
