package dayeleven

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var allMonkeyData = `Monkey 0:
  Starting items: 71, 86
  Operation: new = old * 13
  Test: divisible by 19
    If true: throw to monkey 6
    If false: throw to monkey 7

Monkey 1:
  Starting items: 66, 50, 90, 53, 88, 85
  Operation: new = old + 3
  Test: divisible by 2
    If true: throw to monkey 5
    If false: throw to monkey 4
`

func TestParseStartingItems(t *testing.T) {
	assert.Equal(t, []int{71, 86}, parseStartingItems("  Starting items: 71, 86"))
	assert.Equal(t, []int{66, 50, 90, 53, 88, 85}, parseStartingItems("  Starting items: 66, 50, 90, 53, 88, 85"))
}

func TestParseOperation(t *testing.T) {
	op := parseOperation("  Operation: new = old * 13")
	old := 2
	assert.Equal(t, 26, op(old))

}

func TestOperations(t *testing.T) {
	testCases := []struct {
		opInput string
		old     int
		want    int
	}{
		{
			opInput: "  Operation: new = old * 13",
			old:     2,
			want:    26,
		},
		{
			opInput: "  Operation: new = old * old",
			old:     2,
			want:    4,
		},
		{
			opInput: "  Operation: new = old + 5",
			old:     2,
			want:    7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.opInput, func(t *testing.T) {
			op := parseOperation(tc.opInput)
			assert.Equal(t, tc.want, op(tc.old))
		})
	}
}

func TestParseWorryDivisor(t *testing.T) {
	assert.Equal(t, 2, parseWorryDivisor("  Test: divisible by 2"))
	assert.Equal(t, 5, parseWorryDivisor("  Test: divisible by 5"))
}

func TestGetPassMonkey(t *testing.T) {
	assert.Equal(t, 2, getPassMonkey("    If true: throw to monkey 2"))
}

func TestGetTrueMonkey(t *testing.T) {
	assert.Equal(t, 2, getPassMonkey("    If true: throw to monkey 2"))
	assert.Panics(t, func() { getTrueMonkey("    If false: throw to monkey 2") })
}

func TestGetMonkeyFromData(t *testing.T) {
	r := strings.NewReader(allMonkeyData)
	allMonkeyDataString := parseAllMonkeys(r)

	assert.Len(t, allMonkeyDataString, 2)

	monkeys := getMonkeysFromData(allMonkeyDataString)
	assert.Len(t, monkeys, 2)

	assert.Equal(t, []int{71, 86}, monkeys[0].items)
	assert.Equal(t, []int{66, 50, 90, 53, 88, 85}, monkeys[1].items)
	assert.Equal(t, 5, monkeys[1].calculateWorry(2))
}

func TestReceiveItem(t *testing.T) {
	m := monkey{items: []int{1}}
	m.receiveItem(3)
	assert.Equal(t, []int{1, 3}, m.items)
}

func TestThrowItem(t *testing.T) {
	m := monkey{items: []int{1, 3}}
	item := m.throwRecalculatedItem()
	assert.Equal(t, []int{3}, m.items)
	assert.Equal(t, 1, item)
}

func TestHasItem(t *testing.T) {
	m := monkey{items: []int{1, 3}}
	m.throwRecalculatedItem()
	assert.True(t, m.hasItems())
	m.throwRecalculatedItem()
	assert.False(t, m.hasItems())
}

func TestSolutions(t *testing.T) {
	assert.Equal(t, 88208, solutionA())
	assert.Equal(t, 21115867968, solutionB())
}
