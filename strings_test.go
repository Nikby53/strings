package string_utils

import "testing"

func TestReverse(t *testing.T) {
	var tests = []struct {
		in   string
		exp  string
		name string
	}{
		{"Hello", "olleH", "Reverse"},
		{"", "", "EmptyString"},
		{"🧡💛", "💛🧡", "non-ASCII"},
	}
	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			res := Reverse(e.in)
			if res != e.exp {
				t.Errorf("Error")
			}
		})
	}
}
