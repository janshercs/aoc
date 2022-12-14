package dayfive

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStacks(t *testing.T) {
	f, err := os.Open("./day_five.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// t.Error(GetStacks(f)) // using this for visual checks, i'm lazy to write a test for this
}

func TestGetStackInput(t *testing.T) {
	in := `line1
line2
`
	r := strings.NewReader(in)

	input := getStackInput(r)
	assert.Equal(t, []string{"line1", "line2"}, input)
}

func TestGetNCols(t *testing.T) {
	assert.Equal(t, 9, getNCols(" 1   2   3   4   5   6   7   8   9 "))
}

func TestPop(t *testing.T) {
	stacks := []stack{{'a', 'b'}}
	var popped rune
	stacks[0], popped = stacks[0].pop()
	assert.Equal(t, 'b', popped)
	assert.Equal(t, stack{'a'}, stacks[0])
}

func TestParseInstruction(t *testing.T) {
	instr := parseInstruction("move 3 from 6 to 2")
	assert.Equal(t, instruction{3, 6, 2}, instr)

	instr = parseInstruction("move 10 from 8 to 5")
	assert.Equal(t, instruction{10, 8, 5}, instr)
}

func TestSolutionA(t *testing.T) {
	assert.Equal(t, "BWNCQRMDB", SolutionA())
}

func TestPopMulti(t *testing.T) {
	stacks := []stack{{'a', 'b', 'c', 'd'}}
	modified, popped := stacks[0].popMulti(2)
	assert.Equal(t, stack{'a', 'b'}, modified)
	assert.Equal(t, stack{'c', 'd'}, popped)
}

func TestPushMulti(t *testing.T) {
	s := stack{'a', 'b'}
	n := stack{'c', 'd'}
	assert.Equal(t, stack{'a', 'b', 'c', 'd'}, s.pushMulti(n))
}

func TestSolutionB(t *testing.T) {
	assert.Equal(t, "NHWZCBNBF", SolutionB())
}
