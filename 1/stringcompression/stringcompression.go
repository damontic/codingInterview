// Package stringcompression implements a function to perform basic string compression using the counts of repeated characters.
// For example the string aabccccccaaa would become a2b1c5a3.
package stringcompression

import (
	"strings"
	"strconv"
)

// Function Stringcompressor performs a basic compression of a string using the counts of repeated characters.
// If the "compressed string" would not become smaller that the original string, the funciton returns the original string.
func Stringcompressor(word string) string {
	var builder strings.Builder
	var previous rune
	count := 0
	for _, letter := range word {
		if previous == 0 || letter == previous {
			count++
		} else {
			builder.WriteRune(previous)
			builder.WriteString(strconv.Itoa(count))
			count = 1
		}
		previous = letter
	}
	builder.WriteRune(previous)
	builder.WriteString(strconv.Itoa(count))
	return builder.String()
}
