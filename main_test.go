package main

import "testing"

func TestPing(t *testing.T) {
	got := ping()
	if got != "hoi pong" {
		t.Errorf("ping() =%v; want pong", got)
	}
}
