package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	bdomain "github.com/tumpal2796/shopee/bill/domain"
	bhandler "github.com/tumpal2796/shopee/bill/handler"
	"github.com/tumpal2796/shopee/database"
	thandler "github.com/tumpal2796/shopee/transaction/handler"
	"github.com/tumpal2796/shopee/transaction/resource"
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
