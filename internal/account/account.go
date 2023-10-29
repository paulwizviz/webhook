// Package account contains account related data model and operations
package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/paulwizviz/webhook/internal/dbutil"

	_ "github.com/mattn/go-sqlite3"
)

var (
	EPRootPath          = "/accounts"
	EPPathAcctQueryByID = fmt.Sprintf("%s/", EPRootPath)
)

type queryByIDHandler struct {
	db *sql.DB
}

func (q queryByIDHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)
		return
	}
	accountID := parts[2]

	// We should have a step here to validate accountID
	order, err := dbutil.QueryOrderByID(q.db, accountID)
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

func NewQueryByIDHandler(db *sql.DB) http.Handler {
	return queryByIDHandler{
		db: db,
	}
}
