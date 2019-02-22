package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tumpal2796/sophee/bill/domain"
)

type Response struct {
	StatusCode int         `json:status_code`
	Data       interface{} `json:"data"`
	Error      string      `json:error`
}

type BillInf interface {
	GetMyBill(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type BillImpl struct {
	DomainBill domain.BillInf
}

func NewBill(bill domain.BillInf) BillInf {
	return &BillImpl{
		DomainBill: bill,
	}
}

func (bill *BillImpl) GetMyBill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response Response
	data, err := bill.DomainBill.GetMyBill()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response.StatusCode = 200
	response.Data = data
	response.Error = ""

	result, err := json.Marshal(&response)
	if err != nil {
		w.Write([]byte("error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
	return
}
