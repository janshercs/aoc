package daythirteen

import (
	"bufio"
	"encoding/json"
	"io"
	utils "solutions/aoc_utils"
)

type pair struct {
	lhs, rhs []interface{}
}

func parseInput(r io.Reader) (p []pair) {
	scanner := bufio.NewScanner(r)
	currentPair := pair{}
	for scanner.Scan() {
		if scanner.Text() == utils.EmptyString {
			p = append(p, currentPair)
			currentPair.lhs, currentPair.rhs = nil, nil // reset current pair
			continue
		}
		if currentPair.lhs == nil {
			currentPair.lhs = parseLine(scanner.Text())
		} else {
			currentPair.rhs = parseLine(scanner.Text())
		}

	}
	p = append(p, currentPair)

	return p
}

func parseLine(s string) []interface{} {
	out := []interface{}{}
	json.Unmarshal([]byte(s), &out)
	return out
}
