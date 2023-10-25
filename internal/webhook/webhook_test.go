package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTransactionHandlerMethod(t *testing.T) {

	testcases := []struct {
		Method      string
		Want        int
		description string
	}{
		{
			Method:      http.MethodPost,
			Want:        http.StatusAccepted,
			description: "Method POST",
		},
		{
			Method:      http.MethodGet,
			Want:        http.StatusMethodNotAllowed,
			description: "Method Get",
		},
	}

	for i, tc := range testcases {
		req := httptest.NewRequest(tc.Method, EPPathTransaction, nil)
		w := httptest.NewRecorder()
		TransactionHandler(w, req)
		resp := w.Result()
		if tc.Want != resp.StatusCode {
			t.Fatalf("Want: %v Got: %v Description: %s", tc.Want, resp.StatusCode, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
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
	resp := w.Result()
	b := resp.Body
	defer resp.Body.Close()
	d, _ := io.ReadAll(b)

	expected := TxnPayload{}
	var got TxnPayload
	json.Unmarshal(d, &expected)

	if reflect.DeepEqual(expected, got) {
		t.Fatalf("Want: %v Got: %v ", expected, got)
	}
}
