package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTwo(t *testing.T) {
	score, err := Solution()
	assert.NoError(t, err)
	assert.Equal(t, 13221, score)
}

func TestDayTwoPartTwo(t *testing.T) {
	score, err := SolutionB()
	assert.NoError(t, err)
	assert.Equal(t, 13221, score)
}

func TestMatchScore(t *testing.T) {
	assert.Equal(t, 8, MatchScore("A", "Y"))
	assert.Equal(t, 1, MatchScore("B", "X"))
	assert.Equal(t, 6, MatchScore("C", "Z"))
}

func TestEqualShapes(t *testing.T) {
	x := rockShape{}
	y := rockShape{}
	assert.Equal(t, x, y)
}

func TestSplitLine(t *testing.T) {
	in := splitLine("C Y")
	assert.Equal(t, in[0], "C")
	assert.Equal(t, in[1], "Y")
}
