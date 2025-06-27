package main

import "fmt"

// DistributedTx 定义分布式事务接口
type DistributedTx interface {
	LocalTransaction() bool
	WriteMessageTable()
	AsyncDeliver()
}

// MsgTableTx 实现本地消息表事务
type MsgTableTx struct{}

func (m *MsgTableTx) LocalTransaction() bool {
	fmt.Println("[MsgTable] 执行本地事务")
	return true
}

func (m *MsgTableTx) WriteMessageTable() {
	fmt.Println("[MsgTable] 写入本地消息表")
}

func (m *MsgTableTx) AsyncDeliver() {
	fmt.Println("[MsgTable] 异步投递消息")
}

func main() {
	fmt.Println("MsgTable Demo: Start")
	tx := &MsgTableTx{}
	if tx.LocalTransaction() {
		tx.WriteMessageTable()
		tx.AsyncDeliver()
	}
	fmt.Println("MsgTable Demo: End")
}
