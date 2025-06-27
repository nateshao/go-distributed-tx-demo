package main

import "testing"

func TestStrongDBTx(t *testing.T) {
	tx := &StrongDBTx{}
	if !tx.DistributedTx() {
		t.Error("DistributedTx should return true")
	}
}
