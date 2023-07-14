package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueLetter(t *testing.T) {
	runePtr := func(r rune) *rune {
		return &r
	}

	testCases := []struct {
		text string
		exp  *rune
	}{
		{
			text: "abc",
			exp:  runePtr('a'),
		},
		{
			text: "qq abc",
			exp:  runePtr('a'),
		},
		{
			text: "bb",
			exp:  nil,
		},
		{
			text: `C makes it easy for you to shoot yourself in the foot. C++ makes that harder, but when you do, it blows away your whole leg. (с) Bjarne Stroustrup`,
			exp:  runePtr('e'),
		},
	}

	for i, testCase := range testCases {
		testCase := testCase
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			act := uniqueLetter(testCase.text)
			assert.Equal(t, testCase.exp, act)
		})
	}
}
