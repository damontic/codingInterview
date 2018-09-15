// Package unique implements an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
package unique

// HasAllUniqueCharactersNoDS implements an algorithm to determine if a string has all unique characters without using data structures at all.
func HasAllUniqueCharactersNoDS(word string) bool {
	for _, letter := range word {
		found := 0
		for _, another := range word {
			if letter == another {
				found += 1
				if found > 1 {
					return false
				}
			}
		}
	}
	return true
}

// HasAllUniqueCharacters implements an algorithm to determine if a string has all unique characters using a hashmap.
func HasAllUniqueCharacters(word string) bool {
	counter := make(map[rune]int)
	for _, letter := range word {
		if _, prs := counter[letter]; prs {
			return false
		} else {
			counter[letter] = 1
		}
	}
	return true
}
