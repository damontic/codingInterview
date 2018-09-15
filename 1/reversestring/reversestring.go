// Package reversestring implements a function which reverses a string
package reversestring

import "strings"

func ReverseString(word string) string {
	runes := []rune(word)
	quantity := len(runes)
	var builder strings.Builder
	for i:= quantity-1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	return builder.String()
}
