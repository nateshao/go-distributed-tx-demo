package main

import "fmt"

// TCCDemo 演示TCC事务的基本流程
func main() {
	fmt.Println("TCC Demo: Try-Confirm-Cancel Start")
	if try() {
		confirm()
	} else {
		cancel()
	}
	fmt.Println("TCC Demo: End")
}

func try() bool {
	fmt.Println("[TCC] Try: 资源预留")
	// 模拟资源预留成功
	return true
}

func confirm() {
	fmt.Println("[TCC] Confirm: 正式提交")
}

func cancel() {
	fmt.Println("[TCC] Cancel: 释放资源")
}
