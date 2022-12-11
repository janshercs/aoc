package daythree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitRucksack(t *testing.T) {
	first, second := splitRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	assert.Len(t, first, 12)
	assert.Equal(t, len(first), len(second))
}

func TestFindCommonAlphabet(t *testing.T) {
	alphabet, err := findCommonAlphabet("vJrwpWtwJgWr", "hcsFMMfFFhFp")
	assert.NoError(t, err)
	assert.Equal(t, "p", alphabet)
}

func TestAlphabetToPriority(t *testing.T) {
	assert.Equal(t, 16, alphabetToPriority('p'))
	assert.Equal(t, 38, alphabetToPriority('L'))
}

func TestSolutionA(t *testing.T) {
	i, err := SolutionA()
	assert.NoError(t, err)
	assert.Equal(t, 8053, i)
}

func TestSolutionB(t *testing.T) {
	i, err := SolutionB()
	assert.NoError(t, err)
	assert.Equal(t, 2425, i)
}

func TestIndexRucksack(t *testing.T) {
	assert.Equal(
		t,
		map[rune]bool{
			'a': true,
			'b': true,
			'c': true,
		},
		indexRucksack("abc"))
}
