// Package contains operations related to webhook operations
package webhook

import (
	"encoding/json"
	"fmt"
	"io"
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

	body := r.Body
	defer r.Body.Close()

	if body == nil {
		http.Error(rw, "Empty body requests are not accepted.", http.StatusBadRequest)
		return
	}

	b, err := io.ReadAll(body)
	if err != nil {
		http.Error(rw, "Unable to process body.", http.StatusBadRequest)
		return
	}

	var payLoad TxnPayload
	err = json.Unmarshal(b, &payLoad)
	if err != nil {
		http.Error(rw, "Invalid body.", http.StatusBadRequest)
		return
	}

	// TODO
	// Validate payload
	// Insert to DB it transaction ID are the same it will not commit
	rw.WriteHeader(http.StatusAccepted)
}
