package main

import (
	"fmt"
	"unicode"
)

func firstUniqueLetter(input []rune) rune {
	letterFreq := make(map[rune]int)

	for _, r := range input {
		letterFreq[r]++
	}

	for _, r := range input {
		if letterFreq[r] == 1 {
			return r
		}
	}

	return 0
}

func uniqueLetter(text string) string {

	var runes []rune
	for _, r := range text {
		runes = append(runes, r)
	}

	var uniqueLetters []rune

	start := 0
	for start < len(runes) {
		i := start

		for i < len(runes) && unicode.IsLetter(runes[i]) {
			i++
		}

		uniqueLetter := firstUniqueLetter(runes[start:i])

		if uniqueLetter != 0 {
			uniqueLetters = append(uniqueLetters, uniqueLetter)
		}

		start = i + 1
	}

	return string(firstUniqueLetter(uniqueLetters))
}

func main() {
	text := `C makes it easy for you to shoot yourself in the foot. C++ makes that harder, but when you do, it blows away your whole leg. (с) Bjarne Stroustrup`

	fmt.Println(uniqueLetter(text))
}
