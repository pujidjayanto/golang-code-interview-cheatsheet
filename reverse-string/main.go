package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(ReverseString(s))
	fmt.Println(RecursiveReverseString(s))
}

func ReverseString(s string) string {
	// Convert string to rune slice so Unicode characters are safe
	runes := []rune(s)

	// Reverse runes in-place
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func RecursiveReverseString(s string) string {
	runes := []rune(s)
	reverseHelper(runes, 0, len(runes)-1)
	return string(runes)
}

func reverseHelper(r []rune, left, right int) {
	if left >= right {
		return
	}
	r[left], r[right] = r[right], r[left]
	reverseHelper(r, left+1, right-1)
}
