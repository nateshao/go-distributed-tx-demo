package main

import "fmt"

// SagaDemo 演示SAGA事务的基本流程
func main() {
	fmt.Println("SAGA Demo: Start")
	if step1() && step2() {
		fmt.Println("[SAGA] 所有步骤成功，无需补偿")
	} else {
		compensate2()
		compensate1()
	}
	fmt.Println("SAGA Demo: End")
}

func step1() bool {
	fmt.Println("[SAGA] 步骤1：扣减库存")
	return true
}

func step2() bool {
	fmt.Println("[SAGA] 步骤2：扣减余额")
	// 模拟失败
	return false
}

func compensate2() {
	fmt.Println("[SAGA] 补偿2：返还余额")
}

func compensate1() {
	fmt.Println("[SAGA] 补偿1：返还库存")
}
