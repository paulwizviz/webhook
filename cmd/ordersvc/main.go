package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/paulwizviz/webhook/internal/account"
	"github.com/paulwizviz/webhook/internal/dbutil"
	"github.com/paulwizviz/webhook/internal/hdlutil"
	"github.com/paulwizviz/webhook/internal/webhook"

	_ "github.com/mattn/go-sqlite3"
)

var port uint

func main() {

	// TO DO flags to enable DevOp to specify port

	if port == 0 {
		port = 8080
	}

	db, err := dbutil.ConnectMemDefault()
	if err != nil {
		log.Fatal("Unable to connect to DB")
	}
	defer db.Close()

	err = dbutil.CreateOrderTable(db)
	if errors.Is(err, dbutil.ErrTable) {
		log.Printf("DB table not create. Reason: %v", err)
	}

	txHandler := hdlutil.MethodPostValidate(webhook.TransactionHandler)
	http.HandleFunc(webhook.EPPathTransaction, txHandler)

	acctQureyByIDHandler := hdlutil.MethodGetValidate(account.QueryByIDHandler)
	http.HandleFunc(account.EPPathAcctQueryByID, acctQureyByIDHandler)

	// Operation to start server
	log.Printf("Starting server on: %d", port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}
