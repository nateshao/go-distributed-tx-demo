package main

import "fmt"

// TwoPhaseCommitDemo 演示两阶段提交的基本流程
func main() {
	fmt.Println("2PC Demo: Two-Phase Commit Start")
	// 阶段一：准备
	if prepare() {
		// 阶段二：提交
		commit()
	} else {
		rollback()
	}
	fmt.Println("2PC Demo: End")
}

func prepare() bool {
	fmt.Println("[Coordinator] Prepare phase: All participants ready?")
	// 模拟所有参与者都准备好
	return true
}

func commit() {
	fmt.Println("[Coordinator] Commit phase: All participants commit.")
}

func rollback() {
	fmt.Println("[Coordinator] Rollback phase: Transaction aborted.")
}
