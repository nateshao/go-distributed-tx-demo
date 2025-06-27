package main

import "testing"

func TestTCC(t *testing.T) {
	tx := &TCC{}
	if !tx.Try() {
		t.Error("Try should return true")
	}
	tx.Confirm() // 只需保证不panic
	tx.Cancel()  // 只需保证不panic
}
