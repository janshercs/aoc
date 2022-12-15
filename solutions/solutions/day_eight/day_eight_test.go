package dayeight

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, parseLine("123"))
}

func TestGetGrid(t *testing.T) {
	f, err := os.Open("./day_eight.txt")
	assert.NoError(t, err)
	defer f.Close()

	assert.NotPanics(t, func() { getGrid(f) })
}

func TestGetColsAndRows(t *testing.T) {
	grid := grid{
		{1, 2, 3},
		{1, 2, 3},
	}

	assert.Equal(t, 3, getNCols(grid))
	assert.Equal(t, 2, getNRows(grid))
}

func TestMakeGrid(t *testing.T) {
	grid := grid{
		{0, 0, 0},
		{0, 0, 0},
	}

	assert.Equal(t, grid, makeGrid(3, 2))
}

func TestSolutions(t *testing.T) {
	assert.Equal(t, 1840, solutionA())
	assert.Equal(t, 405769, solutionB())
}

func TestGetRight(t *testing.T) {
	grid := grid{
		{0, 1, 0},
		{1, 0, 1},
	}
	assert.Equal(t, 1, getRightScore(0, 0, grid))
	assert.Equal(t, 2, getRightScore(1, 0, grid))
	assert.Equal(t, 0, getRightScore(1, 2, grid))
	assert.Equal(t, 0, getRightScore(0, 2, grid))
	assert.Equal(t, 1, getLeftScore(0, 2, grid))
	assert.Equal(t, 2, getLeftScore(1, 2, grid))
}
