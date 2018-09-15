// Package replaces all spaces in a string with '%20'.
package urlspaces

import "strings"

func Urlspaces(sentence string) string {
	var builder strings.Builder
	for _, letter := range sentence {
		if letter == rune(' ') {
			builder.WriteString("%20")
		} else {
			builder.WriteRune(letter)
		}
	}
	return builder.String()
}
