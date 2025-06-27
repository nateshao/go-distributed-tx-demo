package main

import "fmt"

// DistributedTx 定义分布式事务接口
type DistributedTx interface {
	SendMsg()
	CheckStatus() bool
}

// ReliableMsgTx 实现可靠消息事务
type ReliableMsgTx struct{}

func (r *ReliableMsgTx) SendMsg() {
	fmt.Println("[ReliableMsg] 发送消息到MQ")
}

func (r *ReliableMsgTx) CheckStatus() bool {
	fmt.Println("[ReliableMsg] 状态回查")
	// 模拟状态回查失败
	return false
}

// ReliableMsgDemo 演示可靠消息+状态检查的基本流程
func main() {
	fmt.Println("ReliableMsg Demo: Start")
	tx := &ReliableMsgTx{}
	tx.SendMsg()
	if tx.CheckStatus() {
		fmt.Println("[ReliableMsg] 消息已成功消费")
	} else {
		fmt.Println("[ReliableMsg] 消息未消费，重试或补偿")
	}
	fmt.Println("ReliableMsg Demo: End")
}
