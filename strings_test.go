package strings

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
				t.Errorf("Expected %q instead of %q", e.exp, res)
			}
		})
	}
}

func TestCount(t *testing.T) {
	var tests = []struct {
		in   string
		exp  int
		name string
	}{
		{"Hello", 5, "Count"},
		{"", 0, "Empty"},
		{"🧡💛®®", 4, "non-ASCII"},
	}
	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			res := Count(e.in)
			if res != e.exp {
				t.Errorf("Expected %v instead of %v", e.exp, res)
			}
		})
	}
}
