package daythirteen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	in := `[1]
[2]

[3]
[4]
`

	assert.Equal(t, []packetPair{
		{[]interface{}{float64(1)}, []interface{}{float64(2)}},
		{[]interface{}{float64(3)}, []interface{}{float64(4)}},
	}, parseInput(strings.NewReader(in)))

}

func TestParseLine(t *testing.T) {
	assert.Equal(t, []interface{}{}, parseLine("[]"))
	assert.Equal(t, []interface{}{float64(1)}, parseLine("[1]"))
	assert.Equal(t, []interface{}{float64(1), float64(32), []interface{}{float64(3)}}, parseLine("[1, 32, [3]]"))
}

// purely developed this based on TDD: I had no idea or inspiration on how to solve this...
func TestInOrder(t *testing.T) {
	testCases := []struct {
		desc string
		pair packetPair
		want bool
	}{
		{
			pair: packetPair{
				lhs: parseLine("[1,1,3,2,2,2,2]"),
				rhs: parseLine("[1,1,5,1,1]"),
			},
			want: true,
		},
		{
			pair: packetPair{
				lhs: parseLine("[[1],[2,3,4]]"),
				rhs: parseLine("[[1],4]"),
			},
			want: true,
		},
		{
			desc: "list in before should pass, and after that should fail.",
			pair: packetPair{
				lhs: parseLine("[[],5]"),
				rhs: parseLine("[[1],4]"),
			},
			want: true,
		},
		{
			desc: "lists in before with same length should pass, and after that should fail.",
			pair: packetPair{
				lhs: parseLine("[[2],5]"),
				rhs: parseLine("[[3],4]"),
			},
			want: true,
		},
		{
			pair: packetPair{
				lhs: parseLine("[7,7,7,7]"),
				rhs: parseLine("[7,7,7]"),
			},
			want: false,
		},
		{
			desc: "test comparing nested lists",
			pair: packetPair{
				lhs: parseLine("[[4,4],4,4]"),
				rhs: parseLine("[[4,4],4,4,4]"),
			},
			want: true,
		},
		{
			desc: "test comparing RHS nested list contains smaller element",
			pair: packetPair{
				lhs: parseLine("[[4,4],4,4]"),
				rhs: parseLine("[[4,3],4,4,4]"),
			},
			want: false,
		},
		{
			desc: "test comparing nested list against int",
			pair: packetPair{
				lhs: parseLine("[9]"),
				rhs: parseLine("[[8,7,6]]"),
			},
			want: false,
		},
		{
			desc: "test different lengths of empty nested lists",
			pair: packetPair{
				lhs: parseLine("[[[]]]"),
				rhs: parseLine("[[]]"),
			},
			want: false,
		},
		{
			desc: "test complex input",
			pair: packetPair{
				lhs: parseLine("[1,[2,[3,[4,[5,6,7]]]],8,9]"),
				rhs: parseLine("[1,[2,[3,[4,[5,6,0]]]],8,9]"),
			},
			want: false,
		},
		{
			desc: "test complex input",
			pair: packetPair{
				lhs: parseLine("[]"),
				rhs: parseLine("[3]"),
			},
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			b, _ := tc.pair.inOrder()
			assert.Equal(t, tc.want, b)
		})
	}
}

func TestSolution(t *testing.T) {
	assert.Equal(t, 5503, solutionA())
	assert.Equal(t, 20952, solutionB())
}
