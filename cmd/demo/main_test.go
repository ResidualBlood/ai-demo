package main

import "testing"

func TestGreeting(t *testing.T) {
	if got := Greeting(""); got != "hello, world" {
		t.Fatalf("got %q", got)
	}
	if got := Greeting("ai"); got != "hello, ai" {
		t.Fatalf("got %q", got)
	}
}
