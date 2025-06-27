package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DistributedTx 定义分布式事务接口
type DistributedTx interface {
	SendMsg()
	CheckStatus() bool
}

// ReliableMsgTx 实现可靠消息事务
type ReliableMsgTx struct {
	db *sql.DB
}

func NewReliableMsgTx(dsn string) (*ReliableMsgTx, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &ReliableMsgTx{db: db}, nil
}

func (r *ReliableMsgTx) SendMsg() {
	_, _ = r.db.Exec("INSERT INTO message_tbl(order_id, status) VALUES(1, 'pending')")
	fmt.Println("[ReliableMsg] 发送消息到MQ（可对接真实MQ）")
}

func (r *ReliableMsgTx) CheckStatus() bool {
	var status string
	err := r.db.QueryRow("SELECT status FROM message_tbl WHERE order_id = 1").Scan(&status)
	if err != nil {
		return false
	}
	return status == "done"
}

// ReliableMsgDemo 演示可靠消息+状态检查的基本流程
func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/db1"
	tx, err := NewReliableMsgTx(dsn)
	if err != nil {
		panic(err)
	}
	tx.SendMsg()
	if tx.CheckStatus() {
		fmt.Println("[ReliableMsg] 消息已成功消费")
	} else {
		fmt.Println("[ReliableMsg] 消息未消费，重试或补偿")
	}
	fmt.Println("ReliableMsg Demo: End")
}
