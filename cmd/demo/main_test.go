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

func TestGreetingTrimsSurroundingSpace(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"  Alice  ", "hello, Alice"},
		{"\tBob\n", "hello, Bob"},
	}
	for _, c := range cases {
		if got := Greeting(c.in); got != c.want {
			t.Fatalf("Greeting(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGreetingBlankBecomesWorld(t *testing.T) {
	for _, in := range []string{"   ", "\t\n"} {
		if got := Greeting(in); got != "hello, world" {
			t.Fatalf("Greeting(%q) = %q, want %q", in, got, "hello, world")
		}
	}
}
