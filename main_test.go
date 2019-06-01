package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hi, Welcome to Shiriff"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
