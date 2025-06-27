package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type StrongDBTx struct {
	db *sql.DB
}

func NewStrongDBTx(dsn string) (*StrongDBTx, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &StrongDBTx{db: db}, nil
}

func (s *StrongDBTx) DistributedTx() bool {
	tx, err := s.db.Begin()
	if err != nil {
		return false
	}
	_, err = tx.Exec("UPDATE account SET balance = balance - 100 WHERE id = 1 AND balance >= 100")
	if err != nil {
		tx.Rollback()
		return false
	}
	return tx.Commit() == nil
}

// StrongDBDemo 演示强一致数据库分布式事务的基本流程
func main() {
	fmt.Println("StrongDB Demo: Start")
	dsn := "user:password@tcp(127.0.0.1:3306)/db1"
	tx, err := NewStrongDBTx(dsn)
	if err != nil {
		panic(err)
	}
	if tx.DistributedTx() {
		fmt.Println("[StrongDB] 分布式事务提交成功")
	} else {
		fmt.Println("[StrongDB] 分布式事务提交失败")
	}
	fmt.Println("StrongDB Demo: End")
}
