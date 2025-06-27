package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TCC struct {
	db *sql.DB
}

func NewTCC(dsn string) (*TCC, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &TCC{db: db}, nil
}

func (t *TCC) Try() bool {
	fmt.Println("[TCC] Try: 资源预留")
	_, err := t.db.Exec("UPDATE inventory SET reserved = reserved + 1 WHERE id = 1 AND stock > reserved")
	return err == nil
}

func (t *TCC) Confirm() {
	fmt.Println("[TCC] Confirm: 正式提交")
	_, _ = t.db.Exec("UPDATE inventory SET stock = stock - 1, reserved = reserved - 1 WHERE id = 1 AND reserved > 0")
}

func (t *TCC) Cancel() {
	fmt.Println("[TCC] Cancel: 释放资源")
	_, _ = t.db.Exec("UPDATE inventory SET reserved = reserved - 1 WHERE id = 1 AND reserved > 0")
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/db1"
	tx, err := NewTCC(dsn)
	if err != nil {
		panic(err)
	}
	if tx.Try() {
		tx.Confirm()
	} else {
		tx.Cancel()
	}
	fmt.Println("TCC Demo: End")
}
