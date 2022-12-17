package dayten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// signal strength = cycle number * register value
func TestParseNoopLine(t *testing.T) {
	ops := parseLine("noop")
	reg := 0
	for _, op := range ops {
		reg = op(reg)
	}
	assert.Equal(t, 0, reg)
}

func TestParseAddxLine(t *testing.T) {
	ops := parseLine("addx 4")
	reg := 0
	for _, op := range ops {
		reg = op(reg)
	}
	assert.Equal(t, 4, reg)

	ops = parseLine("addx -5")
	reg = 0
	for _, op := range ops {
		reg = op(reg)
	}
	assert.Equal(t, -5, reg)
}

func TestSolutions(t *testing.T) {
	assert.Equal(t, 15880, solutionA())
	assert.NotPanics(t, func() { solutionB() })
}

func TestCycleToPos(t *testing.T) {
	assert.Equal(t, 39, getCol(40))
	assert.Equal(t, 0, getCol(1))
	assert.Equal(t, 4, getCol(5))
	assert.Equal(t, 4, getCol(45))
}

func TestPrintRes(t *testing.T) {
	assert.Equal(t, ".", pixelValue(5, 3))
	assert.Equal(t, "#", pixelValue(4, 3))
	assert.Equal(t, ".", pixelValue(4, 0))
	assert.Equal(t, "#", pixelValue(1, 0))
	assert.Equal(t, "#", pixelValue(0, 0))
}

func TestCycleToRow(t *testing.T) {
	assert.Equal(t, 0, getRow(40))
	assert.Equal(t, 1, getRow(41))
}
