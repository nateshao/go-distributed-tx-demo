package main

import "fmt"

// DistributedTx 定义分布式事务接口
type DistributedTx interface {
	Prepare() bool
	Commit()
	Rollback()
}

// TwoPhaseCommit 实现2PC事务
type TwoPhaseCommit struct{}

func (t *TwoPhaseCommit) Prepare() bool {
	fmt.Println("[2PC] Prepare phase: All participants ready?")
	return true
}

func (t *TwoPhaseCommit) Commit() {
	fmt.Println("[2PC] Commit phase: All participants commit.")
}

func (t *TwoPhaseCommit) Rollback() {
	fmt.Println("[2PC] Rollback phase: Transaction aborted.")
}

func main() {
	fmt.Println("2PC Demo: Two-Phase Commit Start")
	tx := &TwoPhaseCommit{}
	if tx.Prepare() {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	fmt.Println("2PC Demo: End")
}
