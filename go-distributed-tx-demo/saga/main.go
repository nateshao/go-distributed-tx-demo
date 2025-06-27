package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SagaTx struct {
	db *sql.DB
}

func NewSagaTx(dsn string) (*SagaTx, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &SagaTx{db: db}, nil
}

func (s *SagaTx) Step1() bool {
	_, err := s.db.Exec("UPDATE inventory SET stock = stock - 1 WHERE id = 1 AND stock > 0")
	return err == nil
}

func (s *SagaTx) Step2() bool {
	_, err := s.db.Exec("UPDATE account SET balance = balance - 100 WHERE id = 1 AND balance >= 100")
	return err == nil
}

func (s *SagaTx) Compensate1() {
	_, _ = s.db.Exec("UPDATE inventory SET stock = stock + 1 WHERE id = 1")
}

func (s *SagaTx) Compensate2() {
	_, _ = s.db.Exec("UPDATE account SET balance = balance + 100 WHERE id = 1")
}

// SagaDemo 演示SAGA事务的基本流程
func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/db1"
	tx, err := NewSagaTx(dsn)
	if err != nil {
		panic(err)
	}
	if tx.Step1() && tx.Step2() {
		fmt.Println("[SAGA] 所有步骤成功，无需补偿")
	} else {
		tx.Compensate2()
		tx.Compensate1()
	}
	fmt.Println("SAGA Demo: End")
}
