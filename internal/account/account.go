// Package account contains account related data model and operations
package account

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/paulwizviz/webhook/internal/dbutil"
)

var (
	EPRootPath          = "/accounts"
	EPPathAcctQueryByID = fmt.Sprintf("%s/", EPRootPath)
)

func QueryByIDHandler(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)
		return
	}
	accountID := parts[2]

	// We should have a step here to validate accountID

	// Ideally we should only make DB connection onces
	db, err := dbutil.ConnectMemDefault()
	if err != nil {
		log.Printf("DB error. %v", err)
	}
	order, err := dbutil.QueryOrderByID(db, accountID)
	if err != nil {
		http.Error(rw, "No data.", http.StatusAccepted)
		return
	}

	ord, err := json.Marshal(order)
	if err != nil {
		http.Error(rw, "No data.", http.StatusAccepted)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
	rw.Write(ord)
}
