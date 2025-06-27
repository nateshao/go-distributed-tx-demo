package main

import "testing"

func TestSagaTx(t *testing.T) {
	tx := &SagaTx{}
	if !tx.Step1() {
		t.Error("Step1 should return true")
	}
	if tx.Step2() {
		t.Error("Step2 should return false (模拟失败)")
	}
	tx.Compensate1()
	tx.Compensate2()
}
