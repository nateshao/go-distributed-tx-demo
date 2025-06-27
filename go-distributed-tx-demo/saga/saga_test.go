package main

import "testing"

func TestStep1(t *testing.T) {
	if !step1() {
		t.Error("step1 should return true")
	}
}

func TestStep2(t *testing.T) {
	if step2() {
		t.Error("step2 should return false (模拟失败)")
	}
}

func TestCompensate(t *testing.T) {
	compensate1()
	compensate2()
}
