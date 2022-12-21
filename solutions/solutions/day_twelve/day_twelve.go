package daytwelve

import (
	"bufio"
	"io"
	"math"
	"os"
	utils "solutions/aoc_utils"
)

// parse input as runes? (or just turn them to int, doesnt matter)
// do bfs for all possible routes without going back
// possible neighbours of a point with value are those with values <= x + 1

type point struct{ x, y int }

func (p point) allNeighbours() []point {
	return []point{
		{p.x + up.x, p.y + up.y},
		{p.x + down.x, p.y + down.y},
		{p.x + left.x, p.y + left.y},
		{p.x + right.x, p.y + right.y},
	}
}

type row []int
type grid []row

const (
	startMarker = 99
	endMarker   = 98
)

var (
	up    = point{-1, 0}
	down  = point{1, 0}
	left  = point{0, -1}
	right = point{0, 1}
)

func getNeighbours(p point, g grid) []point {
	var neighbours []point
	for _, n := range p.allNeighbours() {
		// test in grid
		if !pointInGrid(n, g) {
			continue
		}
		if g[n.x][n.y]-g[p.x][p.y] > 1 {
			continue
		}

		neighbours = append(neighbours, n)
	}

	return neighbours
}

func pointInGrid(p point, g grid) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(g) && p.y < len(g[0])
}

func parseRow(s string) row {
	var r row

	for _, char := range s {
		switch char {
		case 'S':
			r = append(r, 99)
		case 'E':
			r = append(r, 98)
		default:
			i, err := utils.AlphabetPosition(char)
			if err != nil {
				panic(err)
			}
			r = append(r, i)
		}
	}
	return r
}

func readInput(r io.Reader) (g grid, start point, end point) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		g = append(g, parseRow(scanner.Text()))
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if g[i][j] == startMarker {
				start = point{i, j}
				g[i][j] = 1
			}
			if g[i][j] == endMarker {
				end = point{i, j}
				g[i][j] = 26
			}
		}
	}

	return g, start, end
}

func readInputForBetterStart(r io.Reader) (g grid, start []point, end point) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		g = append(g, parseRow(scanner.Text()))
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if g[i][j] == startMarker || g[i][j] == 1 {
				start = append(start, point{i, j})
				g[i][j] = 1
			}
			if g[i][j] == endMarker {
				end = point{i, j}
				g[i][j] = 26
			}
		}
	}

	return g, start, end
}

func bfs(g grid, start point, end point) []point {
	q := queue{{start}}
	visited := map[point]bool{}
	// furthestLength := 15

	for len(q) > 0 {
		currentPath := q[0]
		q = q[1:]

		if !visited[currentPath.lastPoint()] {
			neighbours := getNeighbours(currentPath.lastPoint(), g)
			for _, n := range neighbours {
				if n == end {
					return currentPath
				}
				// newPath := make([]point, 0) // make a new copy
				// newPath = append(newPath, currentPath...)
				// newPath := make([]point, len(currentPath))
				// copy(newPath, currentPath)
				var newPath []point
				newPath = append(newPath, currentPath...)
				newPath = append(newPath, n)
				q = append(q, newPath)
			}

			visited[currentPath.lastPoint()] = true
		}
	}

	return []point{}
}

type queue []path

type path []point

func (p path) lastPoint() point { return p[len(p)-1] }

func solutionA() int {
	f, err := os.Open("./day_twelve.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g, start, end := readInput(f)
	return len(bfs(g, start, end))
}

func solutionB() int {
	f, err := os.Open("./day_twelve.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g, starts, end := readInputForBetterStart(f)
	steps := math.MaxInt

	for _, start := range starts {
		n := len(bfs(g, start, end))
		if n == 0 {
			continue
		}
		if n < steps {
			steps = n
		}

	}
	return steps
}
