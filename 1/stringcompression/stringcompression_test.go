package stringcompression

import "fmt"

func ExampleStringcompressor() {
	fmt.Println(Stringcompressor("aaabbbccc"))
	// Output:
	// a3b3c3
}

func ExampleStringcompressor2() {
	fmt.Println(Stringcompressor("aabcccccaaa"))
	// Output:
	// a2b1c5a3
}
