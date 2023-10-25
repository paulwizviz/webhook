// Package dbutil contains common db operations
package dbutil

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ver = "sqlite3"
)

const (
	transactionId   = "txn_id"
	accountID       = "acct_id"
	orderId         = "order_id"
	transactionType = "txn_type"
	amount          = "amount"
	currency        = "currency"
	description     = "description"
)

var (
	createOrderTableStmtStr = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS order (%s TEXT AS PRIMARY KEY,%s TEXT,%s TEXT,%s TEXT,%s TEXT,%s TEXT,%s TEXT)`, transactionId, accountID, orderId, transactionType, amount, currency, description)
	insertOrderStmtStr      = fmt.Sprintf(`INSERT INTO order (%s,%s,%s,%s,%s,%s,%s,) VALUES ( ?, ?, ?, ?, ?, ?, ? )`, transactionId, accountID, orderId, transactionType, amount, currency, description)
)

var (
	ErrConn         = errors.New("db connect error")
	ErrStmt         = errors.New("statement error")
	ErrInsert       = errors.New("insert error")
	ErrTable        = errors.New("creating table error")
	ErrItemNotFound = errors.New("item not found")
)

type OrderModel struct {
	TxnID       string
	TxnType     string
	OrderID     string
	Amount      string
	Currency    string
	Description string
	AccountID   string
}

func ConnectMemDefault() (*sql.DB, error) {
	db, err := sql.Open(ver, ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

func CreateOrderTable(db *sql.DB) error {
	_, err := db.Exec(createOrderTableStmtStr)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrTable, err.Error())
	}
	return nil
}

func PrepareInsertOrderStmt(db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.Prepare(insertOrderStmtStr)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrStmt, err.Error())
	}
	return stmt, nil
}

func InsertOrder(stmt *sql.Stmt, order OrderModel) (sql.Result, error) {
	result, err := stmt.Exec(order.TxnID, order.AccountID, order.OrderID, order.TxnType, order.Amount, order.Currency, order.Description)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrInsert, err.Error())
	}
	return result, nil
}

func QueryOrderByID(db *sql.DB, id string) (OrderModel, error) {
	row := db.QueryRow("SELECT  transactionId, accountID, orderId, transactionType, amount, currency, description FROM order WHERE accountID=?", id)
	order := OrderModel{}
	err := row.Scan(&order.TxnID, &order.AccountID)
	if err == sql.ErrNoRows {
		log.Printf("Order with id{%v} not found", id)
		return OrderModel{}, fmt.Errorf("%w-%s", ErrItemNotFound, err.Error())
	}
	return order, nil
}
