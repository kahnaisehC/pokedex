package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		in       string
		expected []string
	}{
		{
			"HELLO CHarmaNDER GOOD",
			[]string{"hello", "charmander", "good"},
		},
		{
			"bulS ASAS E??E?",
			[]string{"buls", "asas", "e??e?"},
		},
		{
			"aaaa 131213 3#$$$$",
			[]string{"aaaa", "131213", "3#$$$$"},
		},
	}
	for _, c := range cases {
		in := c.in
		expected := c.expected
		actual := cleanInput(in)

		if len(expected) != len(actual) {
			t.Error("FAIL")
		}
		for i := 0; i < len(expected); i++ {
			if expected[i] != actual[i] {
				t.Errorf("Expected: %s, Actual: %s", expected[i], actual[i])
			}
		}
	}
}
