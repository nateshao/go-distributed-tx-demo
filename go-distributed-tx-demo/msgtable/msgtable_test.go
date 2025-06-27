package main

import "testing"

func TestLocalTransaction(t *testing.T) {
	if !localTransaction() {
		t.Error("localTransaction should return true")
	}
}

func TestWriteMessageTableAndAsyncDeliver(t *testing.T) {
	writeMessageTable() // 只需保证不panic
	asyncDeliver()      // 只需保证不panic
}
