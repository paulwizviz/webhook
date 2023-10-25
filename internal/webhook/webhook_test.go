package webhook

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransactionHandlerEmptyBody(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, EPPathTransaction, nil)
	w := httptest.NewRecorder()
	TransactionHandler(w, req)
	want := http.StatusBadRequest
	got := w.Result().StatusCode

	if want != got {
		t.Fatalf("Want: %d Got %d", want, got)
	}

}

func TestTransactionHandler(t *testing.T) {
	body := `{"transactionId": "tqZi6QapS41zcEHy",
			"transactionType": "SALE",
			"orderId": "c66oxMaisTwJQXjD",
			"amount": "10.00",
			"currency": "EUR",
			"description": "Test transaction",
			"accountId": "001"}`

	req := httptest.NewRequest(http.MethodGet, EPPathTransaction, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	TransactionHandler(w, req)
	want := http.StatusAccepted
	got := w.Result().StatusCode
	if want != got {
		t.Fatalf("Want: %v Got: %v", want, got)
	}
}
