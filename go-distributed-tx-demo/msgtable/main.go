package main

import "fmt"

// MessageTableDemo 演示本地消息表/事务消息的基本流程
func main() {
	fmt.Println("MsgTable Demo: Start")
	if localTransaction() {
		writeMessageTable()
		asyncDeliver()
	}
	fmt.Println("MsgTable Demo: End")
}

func localTransaction() bool {
	fmt.Println("[MsgTable] 执行本地事务")
	// 模拟本地事务成功
	return true
}

func writeMessageTable() {
	fmt.Println("[MsgTable] 写入本地消息表")
}

func asyncDeliver() {
	fmt.Println("[MsgTable] 异步投递消息")
}
