package main

import "testing"

type SagaTx struct{}

func (s *SagaTx) Step1() bool  { return true }
func (s *SagaTx) Step2() bool  { return false }
func (s *SagaTx) Compensate1() {}
func (s *SagaTx) Compensate2() {}

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
