package main

import "fmt"

// ReliableMsgDemo 演示可靠消息+状态检查的基本流程
func main() {
	fmt.Println("ReliableMsg Demo: Start")
	sendMsg()
	if checkStatus() {
		fmt.Println("[ReliableMsg] 消息已成功消费")
	} else {
		fmt.Println("[ReliableMsg] 消息未消费，重试或补偿")
	}
	fmt.Println("ReliableMsg Demo: End")
}

func sendMsg() {
	fmt.Println("[ReliableMsg] 发送消息到MQ")
}

func checkStatus() bool {
	fmt.Println("[ReliableMsg] 状态回查")
	// 模拟状态回查失败
	return false
}
