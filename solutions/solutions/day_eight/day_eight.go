package dayeight

import (
	"bufio"
	"os"
	"strconv"
)

func parseLine(s string) row {
	row := row{}
	for _, r := range s {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		row = append(row, i)
	}
	return row
}

func getGrid(f *os.File) grid {
	grid := grid{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, parseLine(scanner.Text()))
	}
	return grid
}

func getNCols(g grid) int { return len(g[0]) }
func getNRows(g grid) int { return len(g) }

type row []int
type grid []row

func makeGrid(cols, rows int) grid {
	g := grid{}
	for i := 0; i < rows; i++ {
		row := make(row, cols)
		g = append(g, row)
	}
	return g
}

func solutionA() int {
	// get grid
	f, err := os.Open("./day_eight.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	grid := getGrid(f)

	nCols := getNCols(grid)
	nRows := getNRows(grid)

	// make empty grid
	gridMarker := makeGrid(nCols, nRows)

	// start with highest tree 0, keep highest tree encountered so far, if current tree taller than tree, flip marker to 1, highest tree = current tree
	// view from left: iterate left to right via row (j=0; j<NCols; j++)
	for i := 0; i < nRows; i++ {
		highest := 0 // resets every row
		for j := 0; j < nCols; j++ {
			if grid[i][j] > highest {
				gridMarker[i][j] = 1
				highest = grid[i][j]
			}
		}
	}

	// view from right: iterate right to left via row (j=NCols-1; j>= 0; j--)
	for i := 0; i < nRows; i++ {
		highest := 0 // resets every row
		for j := nCols - 1; j >= 0; j-- {
			if grid[i][j] > highest {
				gridMarker[i][j] = 1
				highest = grid[i][j]
			}
		}
	}

	// view from top: iterate top to bottom via col (i:= 0; i<NRows; i++)
	for j := 0; j < nCols; j++ {
		highest := 0 // resets every row
		for i := 0; i < nRows; i++ {
			if grid[i][j] > highest {
				gridMarker[i][j] = 1
				highest = grid[i][j]
			}
		}
	}
	// view from bottom: iterate bottom to top via col (i=NRows-1; i>= 0; i--)
	for j := 0; j < nCols; j++ {
		highest := 0 // resets every row
		for i := nRows - 1; i >= 0; i-- {
			if grid[i][j] > highest {
				gridMarker[i][j] = 1
				highest = grid[i][j]
			}
		}
	}

	visibleTrees := 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if i == 0 || i == nRows-1 || j == 0 || j == nCols-1 { // trees at the edges
				visibleTrees++
				continue
			}
			visibleTrees += gridMarker[i][j]
		}
	}
	return visibleTrees
}

func solutionB() int {
	// get grid
	f, err := os.Open("./day_eight.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	grid := getGrid(f)

	nCols := getNCols(grid)
	nRows := getNRows(grid)

	highest := 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if score := scenicScore(i, j, grid); score > highest {
				highest = score
			}
		}
	}
	return highest
}

// for every tree[i][j], count up,down,left,right until tree >= tree[i][j]

func scenicScore(i, j int, g grid) int {
	return getRightScore(i, j, g) * getLeftScore(i, j, g) * getDownScore(i, j, g) * getUpScore(i, j, g)
}

func getRightScore(i, j int, g grid) int {
	if j == getNCols(g)-1 {
		return 0
	}
	count := 0
	tree := g[i][j]
	for j := j + 1; j < getNCols(g); j++ {
		if g[i][j] < tree {
			count++
			continue
		}
		count++
		break
	}

	return count
}

func getLeftScore(i, j int, g grid) int {
	if j == 0 {
		return 0
	}

	count := 0
	tree := g[i][j]
	for j := j - 1; j >= 0; j-- {
		if g[i][j] < tree {
			count++
			continue
		}
		count++
		break
	}

	return count
}

func getUpScore(i, j int, g grid) int {
	if i == 0 {
		return 0
	}
	count := 0
	tree := g[i][j]
	for i := i - 1; i >= 0; i-- {
		if g[i][j] < tree {
			count++
			continue
		}
		count++
		break
	}

	return count
}

func getDownScore(i, j int, g grid) int {
	nRows := getNRows(g)
	if i == nRows-1 {
		return 0
	}
	count := 0
	tree := g[i][j]
	for i := i + 1; i < nRows; i++ {
		if g[i][j] < tree {
			count++
			continue
		}
		count++
		break
	}

	return count
}
