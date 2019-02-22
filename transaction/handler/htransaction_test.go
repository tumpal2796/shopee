package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/tumpal2796/sophee/transaction/resource/mocks"
)

func TestAddMyBill(t *testing.T) {
	assert := assert.New(t)
	rTransactionMock := new(mocks.TransactionInf)
	hTransaction := NewTransactionHandler(rTransactionMock)
	w := httptest.NewRecorder()
	param := `
	{
		"name": "orer",
		"tax_code": 1,
		"price": 10000
	}`
	r := httptest.NewRequest("POST", "http://localhost:8008/addbill", strings.NewReader(param))
	r.Header.Set("Accept-Language", "en-US")
	p := httprouter.Params{}

	rTransactionMock.On("Insert", mock.Anything, mock.Anything).Return(nil)
	hTransaction.AddMyBill(w, r, p)

	var result Response
	bodyBytes, err := ioutil.ReadAll(w.Body)
	assert.NoError(err, "[TestAddMyBill] Failed")
	err = json.Unmarshal(bodyBytes, &result)
	assert.NoError(err, "[TestAddMyBill] Failed")

	assert.Equal(result.StatusCode, 200, "[TestAddMyBill] Failed")
	t.Logf("[TestAddMyBill] Success")
	return
}
