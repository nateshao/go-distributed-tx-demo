package main

import "testing"

func TestTwoPhaseCommit(t *testing.T) {
	tx := &TwoPhaseCommit{}
	if !tx.Prepare() {
		t.Error("Prepare should return true")
	}
	tx.Commit()   // 只需保证不panic
	tx.Rollback() // 只需保证不panic
}
