package main

import "testing"

func TestReliableMsgTx(t *testing.T) {
	tx := &ReliableMsgTx{}
	tx.SendMsg() // 只需保证不panic
	if tx.CheckStatus() {
		t.Error("CheckStatus should return false (模拟未消费)")
	}
}
