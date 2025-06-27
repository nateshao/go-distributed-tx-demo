package main

import "testing"

type StrongDBTx struct{}

func (s *StrongDBTx) DistributedTx() bool { return true }

func TestStrongDBTx(t *testing.T) {
	tx := &StrongDBTx{}
	if !tx.DistributedTx() {
		t.Error("DistributedTx should return true")
	}
}
