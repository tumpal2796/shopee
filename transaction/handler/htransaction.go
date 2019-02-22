package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sophee/database"
	"github.com/sophee/transaction/model"
	"github.com/sophee/transaction/resource"
)

type Response struct {
	StatusCode int         `json:status_code`
	Data       interface{} `json:"data"`
	Error      string      `json:error`
}

func AddMyBill(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response Response

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	db, err := database.GetPSQLDB()
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	tresource := resource.New(db)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var params model.Transaction
	err = json.Unmarshal([]byte(body), &params)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = tresource.Insert(ctx, params)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response.StatusCode = 200
	response.Error = ""
	response.Data = "Success to Add Bill"

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
