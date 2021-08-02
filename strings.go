package string_utils

// Reverse function reverses a string
func Reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
