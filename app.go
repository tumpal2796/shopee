package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	bhandler "github.com/sophee/bill/handler"
	thandler "github.com/sophee/transaction/handler"
)

func main() {
	router := httprouter.New()
	router.GET("/getmybill", bhandler.GetMyBill)
	router.POST("/addbill", thandler.AddMyBill)

	log.Fatal(http.ListenAndServe(":8080", router))
}
