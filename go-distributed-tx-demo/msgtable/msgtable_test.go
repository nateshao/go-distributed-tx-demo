package main

import "testing"

func TestMsgTableTx(t *testing.T) {
	tx := &MsgTableTx{}
	if !tx.LocalTransaction() {
		t.Error("LocalTransaction should return true")
	}
	tx.WriteMessageTable() // 只需保证不panic
	tx.AsyncDeliver()      // 只需保证不panic
}
