package daythirteen

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"reflect"
	utils "solutions/aoc_utils"
	"sort"
)

type packetPair struct {
	lhs, rhs []interface{}
}

func (p packetPair) inOrder() (bool, bool) {
	shorterLen := shorter(len(p.lhs), len(p.rhs))

	for i, j := 0, 0; i < shorterLen; i, j = i+1, j+1 {
		// find out type of lhs and rhs
		// decide on a way to compare both
		// if int convert to slice and compare slice v.s. slice

		// find out if either are lists first.
		if oneIsList(p.lhs[i], p.rhs[j]) {
			subPacket := packetPair{
				lhs: convertToList(p.lhs[i]),
				rhs: convertToList(p.rhs[i]),
			}

			order, cont := subPacket.inOrder()

			if !order {
				return false, false
			}

			if !cont {
				return true, false
			}

		} else {
			if p.lhs[i].(float64) < p.rhs[j].(float64) {
				return true, false
			}
			if p.lhs[i].(float64) > p.rhs[j].(float64) {
				return false, false
			}
		}
	}

	if len(p.lhs) < len(p.rhs) {
		return true, false
	}

	if len(p.lhs) > len(p.rhs) {
		return false, false
	}

	return true, true
}

func parseInput(r io.Reader) (p []packetPair) {
	scanner := bufio.NewScanner(r)
	currentPair := packetPair{}
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

func shorter(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// returns true if both are slices
func bothAreList(a, b interface{}) bool {
	return reflect.TypeOf(a).Kind() == reflect.Slice && reflect.TypeOf(b).Kind() == reflect.Slice
}

func oneIsList(a, b interface{}) bool {
	return reflect.TypeOf(a).Kind() == reflect.Slice || reflect.TypeOf(b).Kind() == reflect.Slice
}

func convertToList(a interface{}) []interface{} {
	if reflect.TypeOf(a).Kind() == reflect.Slice {
		return a.([]interface{})
	}
	// if not slice, it's integer
	return []interface{}{a}
}

func solutionA() int {
	f, err := os.Open("./day_thirteen.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	packetPairs := parseInput(f)

	indexes := []int{}
	for i := range packetPairs {
		if b, _ := packetPairs[i].inOrder(); b {
			indexes = append(indexes, i)
		}
	}

	sum := len(indexes) // to take care of the 1 based system they want us to sum in.
	for _, i := range indexes {
		sum += i
	}

	return sum
}

// sort all the packets using the above as a sorting algorithm
type packetArray [][]interface{}

func (a packetArray) Len() int      { return len(a) }
func (a packetArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a packetArray) Less(i, j int) bool {
	p := packetPair{a[i], a[j]}
	b, _ := p.inOrder()
	return b
}

func parseAllPackets(r io.Reader) (p packetArray) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Text() == utils.EmptyString {
			continue
		}
		p = append(p, parseLine(scanner.Text()))
	}
	return p
}

func solutionB() int {
	f, err := os.Open("./day_thirteen.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pArray := parseAllPackets(f)
	pArray = append(pArray, parseLine("[[6]]"))
	pArray = append(pArray, parseLine("[[2]]"))
	sort.Sort(pArray)

	decoderIndex := []int{}
	for i, p := range pArray {
		if len(p) != 1 || reflect.TypeOf(p[0]).Kind() != reflect.Slice {
			continue
		}

		pSlice := p[0].([]interface{}) // have to force it to []interface{} to compare the final interface{} that is inside

		if len(pSlice) != 1 {
			continue
		}

		if pSlice[0].(float64) == float64(6) || pSlice[0].(float64) == float64(2) {
			decoderIndex = append(decoderIndex, i+1)
		}
	}
	return decoderIndex[0] * decoderIndex[1]
}
