package daythree

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	utils "solutions/aoc_utils"
	"unicode"
)

var ErrAlphabetNotFound = errors.New("alphabet not found")

func splitRucksack(s string) (string, string) {
	half := len(s) / 2
	return s[:half], s[half:]
}

func findCommonAlphabet(first, second string) (rune, error) {
	firstSet := map[rune]bool{}
	for _, r := range first {
		firstSet[r] = true
	}

	for _, needle := range second {
		if firstSet[needle] {
			return needle, nil
		}
	}
	return ' ', ErrAlphabetNotFound
}

func alphabetToPriority(in rune) int {
	priority, _ := utils.AlphabetPosition(in)

	casePriority := 0
	isUpper := in == unicode.ToUpper(in)
	if isUpper {
		casePriority = 26
	}

	return priority + casePriority
}

func SolutionA() (int, error) {
	f, err := os.Open("./day_three.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalPriority := 0

	for scanner.Scan() {
		alphabet, err := findCommonAlphabet(splitRucksack(scanner.Text()))
		if err != nil {
			return 0, fmt.Errorf("failed to find common letter: %w", err)
		}
		totalPriority += alphabetToPriority(alphabet)
	}
	return totalPriority, nil
}

func SolutionB() (int, error) {
	f, err := os.Open("./day_three.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPriority := 0
	rucksacks := []rucksack{}
	for scanner.Scan() {
		rucksacks = append(rucksacks, indexRucksack(scanner.Text()))
	}

	for i := 0; i <= len(rucksacks)-3; i += 3 {
		badge, err := findCommonBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		if err != nil {
			return 0, fmt.Errorf("unable to check if item is in rucksack: %w", err)
		}
		totalPriority += alphabetToPriority(badge)
	}

	return totalPriority, nil

}

func indexRucksack(s string) rucksack {
	contents := rucksack{}
	for _, r := range s {
		contents[r] = true
	}
	return contents
}

func findCommonBadge(first, second, third rucksack) (rune, error) {
	possibleRunes := []rune{}
	for r := range first {
		if second[r] {
			possibleRunes = append(possibleRunes, r)
		}
	}

	for _, r := range possibleRunes {
		if third[r] {
			return r, nil
		}
	}
	return ' ', fmt.Errorf("failed to find common letter")
}

type rucksack map[rune]bool
