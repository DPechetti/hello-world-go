package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Donovan"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello (%q) = %q, %v, want match for %q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)

	if msg != "" || err == nil {
		t.Fatalf(`Hello (%q) = %q, %v, want %q, nil`, name, msg, err, "")
	}
}
