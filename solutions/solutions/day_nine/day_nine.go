package daynine

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// no need for grid? no idea how much it is going to walk to.
// can keep position of head and tail with i,j
// check touching (include diagonal)
// check direction of tail to follow

const NKnots = 10

var (
	up    = point{0, 1}
	down  = point{0, -1}
	left  = point{-1, 0}
	right = point{1, 0}
)

type point struct{ x, y int }

func (p point) move(q point) point { return point{x: p.x + q.x, y: p.y + q.y} }

func isTouching(head, tail point) bool {
	if abs(head.x-tail.x) > 1 {
		return false
	}
	if abs(head.y-tail.y) > 1 {
		return false
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func isSameRow(head, tail point) bool {
	return head.x == tail.x
}

func isSameCol(head, tail point) bool {
	return head.y == tail.y
}

func headIsRight(head, tail point) bool {
	return head.x-tail.x > 0
}
func headIsUp(head, tail point) bool {
	return head.y-tail.y > 0
}

func getTailDisplacement(head, tail point) point {
	if isSameCol(head, tail) {
		if head.x-tail.x > 0 {
			return point{1, 0}
		}
		return point{-1, 0}
	}

	if isSameRow(head, tail) {
		if head.y-tail.y > 0 {
			return point{0, 1}
		}
		return point{0, -1}
	}

	x, y := -1, -1

	if headIsRight(head, tail) {
		x = 1
	}
	if headIsUp(head, tail) {
		y = 1
	}

	// diagonal right up: {1, 1}
	// diagonal right down: {1, -1}
	// diagonal left up: {-1, 1}
	// diagonal left down: {-1, -1}

	return point{x, y}
}

type instruction struct {
	direction point
	qty       int
}

func parseLine(s string) instruction {
	in := strings.Split(s, " ")
	qty, err := strconv.Atoi(in[1])
	if err != nil {
		panic(err)
	}
	var dir point
	switch in[0] {
	case "U":
		dir = up
	case "D":
		dir = down
	case "L":
		dir = left
	case "R":
		dir = right
	default:
		panic("did not match any direction from input")
	}

	return instruction{direction: dir, qty: qty}
}

func solutionA() int {
	f, err := os.Open("./day_nine.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	head, tail := point{0, 0}, point{0, 0}
	scanner := bufio.NewScanner(f)

	visited := map[point]bool{{0, 0}: true} // starting position
	for scanner.Scan() {
		instr := parseLine(scanner.Text())
		for i := 0; i < instr.qty; i++ {
			head = head.move(instr.direction)
			if isTouching(head, tail) {
				continue
			}

			tail = tail.move(getTailDisplacement(head, tail))
			visited[tail] = true
		}
	}

	return len(visited)
}

func solutionB() int {
	f, err := os.Open("./day_nine.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	head := 0
	tail := NKnots - 1
	scanner := bufio.NewScanner(f)
	knotPosition := map[int]point{}
	for i := 0; i < NKnots; i++ {
		knotPosition[0] = point{0, 0}
	}

	visited := map[point]bool{{0, 0}: true} // starting position
	for scanner.Scan() {
		instr := parseLine(scanner.Text())

		for i := 0; i < instr.qty; i++ {
			knotPosition[head] = knotPosition[head].move(instr.direction) // only move the head
			for k := 0; k < NKnots-1; k++ {
				if isTouching(knotPosition[k], knotPosition[k+1]) {
					continue
				}
				knotPosition[k+1] = knotPosition[k+1].move(getTailDisplacement(knotPosition[k], knotPosition[k+1]))

				if k+1 == tail {
					visited[knotPosition[tail]] = true
				}
			}
		}
	}

	return len(visited)
}
