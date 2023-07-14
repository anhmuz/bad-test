package main

import (
	"fmt"
	"unicode"
)

func firstUniqueLetter(input []rune) *rune {
	letterFreq := make(map[rune]int)

	for _, r := range input {
		letterFreq[r]++
	}

	for _, r := range input {
		if letterFreq[r] == 1 {
			return &r
		}
	}

	return nil
}

func uniqueLetter(text string) *rune {
	runes := ([]rune)(text)

	var uniqueLetters []rune

	start := 0
	for start < len(runes) {
		i := start

		for i < len(runes) && !unicode.IsSpace(runes[i]) {
			i++
		}

		uniqueLetter := firstUniqueLetter(runes[start:i])

		if uniqueLetter != nil {
			uniqueLetters = append(uniqueLetters, *uniqueLetter)
		}

		start = i + 1
	}

	return firstUniqueLetter(uniqueLetters)
}

func main() {
	text := `C makes it easy for you to shoot yourself in the foot. C++ makes that harder, but when you do, it blows away your whole leg. (с) Bjarne Stroustrup`

	r := uniqueLetter(text)

	if r != nil {
		fmt.Printf("%c", *r)
	} else {
		fmt.Println("not found")
	}
}
