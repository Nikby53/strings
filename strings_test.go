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
				t.Errorf("Expected %v instead of %v", e.exp, e.in)
			}
		})
	}
}

func TestCount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		word := "Hello"
		expected := 5
		actual := Count(word)
		if actual != expected {
			t.Errorf("Expected %v instead of %v", expected, actual)
		}
	})
}
