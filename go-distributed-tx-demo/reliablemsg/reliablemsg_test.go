package main

import "testing"

func TestSendMsg(t *testing.T) {
	sendMsg() // 只需保证不panic
}

func TestCheckStatus(t *testing.T) {
	if checkStatus() {
		t.Error("checkStatus should return false (模拟未消费)")
	}
}
