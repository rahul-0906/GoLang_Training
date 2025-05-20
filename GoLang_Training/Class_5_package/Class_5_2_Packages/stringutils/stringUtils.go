package stringutils

import "strings"

// Reverse
func Reverse(s string) string {
	runes := []rune(s)

	i, j := 0, len(s)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

// ToUpperCase

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

func CountVowels(s string) int {
	count := 0

	for _, val := range strings.ToLower(s) {
		switch val {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
