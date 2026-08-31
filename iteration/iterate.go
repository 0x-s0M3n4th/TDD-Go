package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder // initializing an empty string
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// Comparison
func Compare(value_1, value_2 string) bool {
	if value_1 == value_2 {
		return true
	}
	return false
}
