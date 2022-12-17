package dayten

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const addx = "addx"

var keyCycles = map[int]bool{
	20:  true,
	60:  true,
	100: true,
	140: true,
	180: true,
	220: true,
}

func noop(x int) int {
	return x
}

func makeAddx(x int) op {
	return func(i int) int { return i + x }
}

type op func(int) int

func parseLine(s string) []op {
	ops := []op{noop}
	if strings.HasPrefix(s, addx) {
		in := strings.Split(s, " ")
		i, err := strconv.Atoi(in[1])
		if err != nil {
			panic(err)
		}
		ops = append(ops, makeAddx(i))
	}

	return ops
}

// check reg value in key cycles BEFORE doing op.

func solutionA() int {
	f, err := os.Open("./day_ten.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var reg, signal, cycle = 1, 0, 1

	for scanner.Scan() {
		ops := parseLine(scanner.Text())
		for _, op := range ops {
			if keyCycles[cycle] {
				signal += (cycle * reg)
			}
			reg = op(reg)
			cycle++
		}
	}

	return signal
}

// Solution B
// X represents middle position of a 3 pixel sprite ###
// If the sprite is positioned such that one of its three pixels is the pixel currently being drawn,
// the screen produces a lit pixel (#)

// left most position is 0 and right most is 39, each row has 40 pixels
// for every cycle, check which position in the row the cycle is cycle 1 = pixel 0, cycle 40 = pixel 39, cycle 41 = pixel 0
// check reg value if reg value - 1<= pixel position >= reg value + 1, paint [#] else paint [.]

func getCol(i int) int {
	i = (i - 1) % 40
	return i
}

func getRow(i int) int {
	return (i - 1) / 40
}

func pixelValue(pos, reg int) string {
	if reg-1 <= pos && pos <= reg+1 {
		return "#"
	}
	return "."
}

func solutionB() string {
	monitor := [][]string{}
	for i := 0; i < 6; i++ {
		monitor = append(monitor, make([]string, 40))
	}

	f, err := os.Open("./day_ten.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var reg, cycle = 1, 1

	for scanner.Scan() {
		ops := parseLine(scanner.Text())
		for _, op := range ops {
			i, j := getRow(cycle), getCol(cycle)
			monitor[i][j] = pixelValue(j, reg)

			reg = op(reg)
			cycle++
		}
	}

	for _, line := range monitor {
		fmt.Println(line)
	}
	return "hi"
}
