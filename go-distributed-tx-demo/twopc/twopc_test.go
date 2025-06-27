package main

import "testing"

func TestPrepare(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"all ready", true},
	}
	for _, c := range cases {
		got := prepare()
		if got != c.want {
			t.Errorf("%s: want %v, got %v", c.name, c.want, got)
		}
	}
}

func TestCommitAndRollback(t *testing.T) {
	commit()   // 只需保证不panic
	rollback() // 只需保证不panic
}
