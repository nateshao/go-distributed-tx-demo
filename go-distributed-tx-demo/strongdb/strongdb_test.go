package main

import "testing"

func TestDistributedTx(t *testing.T) {
	if !distributedTx() {
		t.Error("distributedTx should return true")
	}
}
