// Package ispermutation defines a function that decides if a string is a permutation of another
package ispermutation

// Function IsPermutation returns true if the string first is a permutation of string second.
func IsPermutation(first, second string) bool {

	if len(first) != len(second) {
		return false
	}

	firstMap := make(map[rune]int)
	secondMap := make(map[rune]int)

	for _, letter := range first {
		if _, prs := firstMap[letter]; prs {
			firstMap[letter] += 1
		} else {
			firstMap[letter] = 1
		}
	}

        for _, letter := range second {
                if _, prs := secondMap[letter]; prs {
                        secondMap[letter] += 1
                } else {
                        secondMap[letter] = 1
                }
        }

	for key, value := range firstMap {
		if elem, prs := secondMap[key]; !prs || elem != value {
			return false
		}
	}

	return true
}
