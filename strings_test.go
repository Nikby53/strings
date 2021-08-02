package string_utils

import "testing"

func TestReverse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		word := "Hello"
		expected := "olleH"
		actual := Reverse(word)
		if actual != expected {
			t.Errorf("Expected %v instead of %v", expected, actual)
		}
	})

}
