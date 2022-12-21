package daytwelve

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNeighbours(t *testing.T) {
	t.Parallel()
	grid := grid{
		{1, 4, 3},
		{1, 0, 3},
		{1, 2, 3},
	}

	testCases := []struct {
		p    point
		want []point
	}{
		{
			p:    point{1, 1},
			want: []point{{1, 0}},
		},
		{
			p:    point{0, 0},
			want: []point{{1, 0}},
		},
		{
			p:    point{0, 1},
			want: []point{{0, 0}, {0, 2}, {1, 1}},
		},
		{
			p:    point{2, 0},
			want: []point{{1, 0}, {2, 1}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run("", func(t *testing.T) {
			ns := getNeighbours(tc.p, grid)
			assert.ElementsMatch(t, tc.want, ns)
		})
	}

}

func TestParseRow(t *testing.T) {
	assert.Equal(t, row{1, 2, 3, 4}, parseRow("abcd"))
	assert.Equal(t, row{99, 2, 26, 98}, parseRow("SbzE"))
}

func TestReadInput(t *testing.T) {
	in := `abc
SaE
abc`
	g, start, end := readInput(strings.NewReader(in))
	assert.Equal(t, point{1, 0}, start)
	assert.Equal(t, point{1, 2}, end)
	assert.Equal(t, grid{
		{1, 2, 3},
		{1, 1, 26},
		{1, 2, 3},
	}, g)
}

func TestSolutions(t *testing.T) {
	assert.Equal(t, 472, solutionA())
	assert.Equal(t, 465, solutionB())
}

func TestBFS(t *testing.T) {
	in := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	assert.Len(t, bfs(readInput(strings.NewReader(in))), 31)
}
