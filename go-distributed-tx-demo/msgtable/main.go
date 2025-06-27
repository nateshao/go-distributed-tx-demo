package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DistributedTx 定义分布式事务接口
type DistributedTx interface {
	LocalTransaction() bool
	WriteMessageTable()
	AsyncDeliver()
}

// MsgTableTx 实现本地消息表事务
type MsgTableTx struct {
	db *sql.DB
}

func NewMsgTableTx(dsn string) (*MsgTableTx, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MsgTableTx{db: db}, nil
}

func (m *MsgTableTx) LocalTransaction() bool {
	tx, err := m.db.Begin()
	if err != nil {
		return false
	}
	_, err = tx.Exec("UPDATE order_tbl SET status = 'created' WHERE id = 1")
	if err != nil {
		tx.Rollback()
		return false
	}
	return tx.Commit() == nil
}

func (m *MsgTableTx) WriteMessageTable() {
	_, _ = m.db.Exec("INSERT INTO message_tbl(order_id, status) VALUES(1, 'pending')")
}

func (m *MsgTableTx) AsyncDeliver() {
	fmt.Println("[MsgTable] 异步投递消息（可对接MQ）")
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/db1"
	tx, err := NewMsgTableTx(dsn)
	if err != nil {
		panic(err)
	}
	if tx.LocalTransaction() {
		tx.WriteMessageTable()
		tx.AsyncDeliver()
	}
	fmt.Println("MsgTable Demo: End")
}
