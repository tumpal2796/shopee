package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sophee/bill/domain"
	"github.com/sophee/database"
	"github.com/sophee/transaction/resource"
)

type Response struct {
	StatusCode int         `json:status_code`
	Data       interface{} `json:"data"`
	Error      string      `json:error`
}

func GetMyBill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response Response

	db, err := database.GetPSQLDB()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tresource := resource.New(db)
	dbill := domain.NewBill(tresource)
	data, err := dbill.GetMyBill()
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
