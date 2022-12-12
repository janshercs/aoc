package dayfour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// if start is smaller or equal to and end is bigger or equal to, then they overlap
// get the smaller start
// if end is bigger or equal to than other then they overlap

func TestGetAssignments(t *testing.T) {
	assignments := getAssignmentPairs("17-38,18-36")
	assert.Equal(t, []assignment{
		{
			start: 17,
			end:   38,
		},
		{
			start: 18,
			end:   36,
		},
	}, assignments)
}

func TestGetSmallerStart(t *testing.T) {
	smaller, bigger := sortByStart([]assignment{{17, 38}, {18, 36}})
	assert.Equal(t, assignment{17, 38}, smaller)
	assert.Equal(t, assignment{18, 36}, bigger)

	smaller, bigger = sortByStart([]assignment{{17, 38}, {17, 37}})
	assert.Equal(t, assignment{17, 38}, smaller)
	assert.Equal(t, assignment{17, 37}, bigger)
}

func TestIsContaining(t *testing.T) {
	assert.Equal(t, true, isContaining(assignment{17, 38}, assignment{18, 36}))
	assert.Equal(t, true, isContaining(assignment{17, 38}, assignment{17, 37}))
	assert.Equal(t, true, isContaining(assignment{17, 38}, assignment{17, 38}))
	assert.Equal(t, true, isContaining(assignment{17, 38}, assignment{17, 19}))
	assert.Equal(t, false, isContaining(assignment{17, 38}, assignment{18, 39}))
}

func TestIsOverlapping(t *testing.T) {
	assert.Equal(t, true, isOverlapping(assignment{17, 38}, assignment{18, 39}))
	assert.Equal(t, false, isOverlapping(assignment{17, 38}, assignment{39, 39}))
}

func TestSolutionA(t *testing.T) {
	assert.Equal(t, 413, SolutionA())
}

func TestSolutionB(t *testing.T) {
	assert.Equal(t, 806, SolutionB())
}
