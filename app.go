package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	bdomain "github.com/tumpal2796/sophee/bill/domain"
	bhandler "github.com/tumpal2796/sophee/bill/handler"
	"github.com/tumpal2796/sophee/database"
	thandler "github.com/tumpal2796/sophee/transaction/handler"
	"github.com/tumpal2796/sophee/transaction/resource"
)

func main() {
	router := httprouter.New()
	db, err := database.GetPSQLDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tresource := resource.New(db)
	thandler := thandler.NewTransactionHandler(tresource)
	router.POST("/addbill", thandler.AddMyBill)

	bdomain := bdomain.NewBill(tresource)
	bhandler := bhandler.NewBill(bdomain)
	router.GET("/getmybill", bhandler.GetMyBill)

	log.Fatal(http.ListenAndServe(":8080", router))
}
