package dayeleven

import (
	"bufio"
	"fmt"
	"io"
	"os"
	utils "solutions/aoc_utils"
	"sort"
	"strconv"
	"strings"
)

// inputs represent a set of rules and start state
// seems like good idea to just represent them as monkeys

type monkey struct {
	items          []int
	trueMonkey     int
	falseMonkey    int
	worryDivisor   int
	calculateWorry func(int) int
}

func (m *monkey) receiveItem(i int) {
	m.items = append(m.items, i)
}

func (m *monkey) throwRecalculatedItem() int {
	item := m.items[0]
	m.items = m.items[1:]
	return m.calculateWorry(item) / 3
}

// there's some math to understand here about
// "You'll need to find another way to keep your worry levels manageable"
// https://www.reddit.com/r/adventofcode/comments/zih7gf/2022_day_11_part_2_what_does_it_mean_find_another/
func (m *monkey) throwItemWithoutRelief(lcm int) int {
	item := m.items[0]
	m.items = m.items[1:]
	return m.calculateWorry(item) % lcm
}

func (m *monkey) hasItems() bool {
	return len(m.items) > 0
}

func (m *monkey) getMonkeyToThrowTo(i int) int {
	if i%m.worryDivisor == 0 {
		return m.trueMonkey
	}
	return m.falseMonkey
}

func parseStartingItems(s string) []int {
	var items []int

	input := strings.Split(s, ":")
	itemString := input[1]
	itemStrings := strings.Split(itemString, ",")

	for _, s := range itemStrings {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			panic(err)
		}
		items = append(items, i)
	}

	return items
}

// only handles * and + because test cases don't contain other ops
func parseOperation(s string) func(int) int {
	//  Operation: new = old * 13
	//  Operation: new = old + 3
	//  Operation: new = old * old

	expression := strings.Split(s, "=")[1] // we only care about the LHS expression
	expression = strings.TrimSpace(expression)
	expressionTerms := strings.Split(expression, " ")
	mathOperation, term := expressionTerms[1], expressionTerms[2]
	switch mathOperation {
	case "+":
		addend, err := strconv.Atoi(term)
		if err != nil {
			panic(err)
		}
		return func(i int) int { return i + addend }
	case "*":
		addend, err := strconv.Atoi(term)
		if err != nil {
			return func(i int) int { return i * i }
		}
		return func(i int) int { return i * addend }
	default:
		panic(fmt.Sprintf("unable to parse operation: %s", s))
	}

}

func parseWorryDivisor(s string) int {
	return utils.GetLastIntInString(s)
}

func getPassMonkey(s string) int {
	return utils.GetLastIntInString(s)
}

func getTrueMonkey(s string) int {
	if strings.HasPrefix(s, "    If true") {
		return getPassMonkey(s)
	}
	panic("wrong statement passed in to get monkey to pass to if true")
}

func getFalseMonkey(s string) int {
	if strings.HasPrefix(s, "    If false") {
		return getPassMonkey(s)
	}
	panic("wrong statement passed in to get monkey to pass to if false")
}

func parseAllMonkeys(r io.Reader) [][]string {
	scanner := bufio.NewScanner(r)
	var input []string
	var monkeys [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == utils.EmptyString {
			monkeys = append(monkeys, input)
			input = []string{} // reset input
			continue
		}
		input = append(input, line)
	}
	monkeys = append(monkeys, input) // last line
	return monkeys
}

func getMonkeysFromData(data [][]string) []monkey {
	var monkeys []monkey
	for _, m := range data {
		monkey := monkey{
			items:          parseStartingItems(m[1]),
			trueMonkey:     getTrueMonkey(m[4]),
			falseMonkey:    getFalseMonkey(m[5]),
			worryDivisor:   parseWorryDivisor(m[3]),
			calculateWorry: parseOperation(m[2]),
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func solutionA() int {
	f, err := os.Open("./day_eleven.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	monkeys := getMonkeysFromData(parseAllMonkeys(f))
	monkeyBusiness := make([]int, len(monkeys))
	var rounds = 20
	for i := 0; i < rounds; i++ {
		for monkeyNo := range monkeys {
			for monkeys[monkeyNo].hasItems() {
				monkeyBusiness[monkeyNo]++
				item := monkeys[monkeyNo].throwRecalculatedItem()
				monkeys[monkeys[monkeyNo].getMonkeyToThrowTo(item)].receiveItem(item)
			}
		}
	}

	sorted := sort.IntSlice(monkeyBusiness)
	sorted.Sort()
	return sorted[len(sorted)-1] * sorted[len(sorted)-2]
}

func solutionB() int {
	f, err := os.Open("./day_eleven.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	monkeys := getMonkeysFromData(parseAllMonkeys(f))
	monkeyBusiness := make([]int, len(monkeys))
	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.worryDivisor
	}

	var rounds = 10000
	for i := 0; i < rounds; i++ {
		for monkeyNo := range monkeys {
			for monkeys[monkeyNo].hasItems() {
				monkeyBusiness[monkeyNo]++
				item := monkeys[monkeyNo].throwItemWithoutRelief(lcm)
				monkeys[monkeys[monkeyNo].getMonkeyToThrowTo(item)].receiveItem(item)
			}
		}
	}

	sorted := sort.IntSlice(monkeyBusiness)
	sorted.Sort()
	return sorted[len(sorted)-1] * sorted[len(sorted)-2]
}
