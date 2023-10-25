// Package contains operations related to webhook operations
package webhook

import (
	"fmt"
	"net/http"
)

var (
	EPRootPath        = "/webhooks"
	EPPathTransaction = fmt.Sprintf("%s/transaction", EPRootPath)
)

type TxnPayload struct {
	ID          string `json:"transactionId"`
	Type        string `json:"transactionType"`
	OrderID     string `json:"orderID"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	AccountID   string `json:"accountId"`
}

func TransactionHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request from %s\n", r.Method, r.URL)

	if r.Method != http.MethodPost {
		http.Error(rw, "Invalid HTTP method. Only POST requests are accepted.", http.StatusMethodNotAllowed)
		return
	}

	rw.WriteHeader(http.StatusAccepted)

}
