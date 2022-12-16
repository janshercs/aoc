package daynine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTouching(t *testing.T) {
	assert.True(t, isTouching(point{1, 1}, point{1, 1}))
	assert.True(t, isTouching(point{1, 1}, point{2, 2}))
	assert.False(t, isTouching(point{1, 1}, point{3, 1}))
	assert.False(t, isTouching(point{1, 1}, point{3, 3}))
}

func TestIsSameRow(t *testing.T) {
	assert.True(t, isSameRow(point{1, 1}, point{1, 3}))
	assert.False(t, isSameRow(point{1, 1}, point{2, 1}))
}

func TestIsSameCol(t *testing.T) {
	assert.True(t, isSameCol(point{2, 3}, point{1, 3}))
	assert.False(t, isSameCol(point{1, 3}, point{2, 1}))
}

func TestGetDisplacement(t *testing.T) {
	assert.Equal(t, point{-1, 0}, getTailDisplacement(point{1, 0}, point{3, 0}))
	assert.Equal(t, point{1, 0}, getTailDisplacement(point{3, 0}, point{1, 0}))
	assert.Equal(t, point{0, 1}, getTailDisplacement(point{0, 3}, point{0, 1}))
	assert.Equal(t, point{0, -1}, getTailDisplacement(point{0, 1}, point{0, 3}))
	assert.Equal(t, point{1, 1}, getTailDisplacement(point{2, 4}, point{1, 1}))
	assert.Equal(t, point{-1, -1}, getTailDisplacement(point{1, 1}, point{2, 4}))
	assert.Equal(t, point{1, -1}, getTailDisplacement(point{2, 1}, point{0, 4}))
	assert.Equal(t, point{-1, 1}, getTailDisplacement(point{0, 3}, point{2, 1}))
}

func TestParseline(t *testing.T) {
	testCases := []struct {
		input string
		instr instruction
	}{
		{
			input: "U 2",
			instr: instruction{up, 2},
		},
		{
			input: "D 5",
			instr: instruction{down, 5},
		},
		{
			input: "L 3",
			instr: instruction{left, 3},
		},
		{
			input: "R 13",
			instr: instruction{right, 13},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.instr, parseLine(tc.input))
		})
	}
}

func TestSolutions(t *testing.T) {
	assert.Equal(t, 6011, solutionA())
	assert.Equal(t, 2419, solutionB())
}
