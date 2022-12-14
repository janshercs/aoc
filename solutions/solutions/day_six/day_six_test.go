package daysix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// start of packet marker = 4 different characters
// double pointer stop when leading - trailing == 4
// need to keep track of characters
// need to remove trailing char from character

func TestMarkerSearch(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "something",
			want:  4,
		},
		{
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
		{
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {

			assert.Equal(t, tc.want, getMarkerPosition(tc.input, 4))
		})
	}
}

func TestSolutionA(t *testing.T) {
	assert.Equal(t, 1655, solutionA())
}

func TestSolutionB(t *testing.T) {
	assert.Equal(t, 2665, solutionB())
}
