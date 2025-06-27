package main

import "testing"

func TestTry(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"reserve success", true},
	}
	for _, c := range cases {
		got := try()
		if got != c.want {
			t.Errorf("%s: want %v, got %v", c.name, c.want, got)
		}
	}
}

func TestConfirmAndCancel(t *testing.T) {
	confirm() // 只需保证不panic
	cancel()  // 只需保证不panic
}
