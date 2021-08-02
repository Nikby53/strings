package string_utils

// Reverse function reverses a string
func Reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

// Count function returns a number of symbols in the string
func Count(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
