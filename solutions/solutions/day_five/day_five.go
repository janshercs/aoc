package dayfive

import (
	"bufio"
	"io"
	"os"
	utils "solutions/aoc_utils"
	"strconv"
	"strings"
)

// gonna need a stack data structure
// gonna need to know how to parse it to initiate the DS
const (
	colWidth    = 4
	runeOffset  = 1
	stackOffset = 1
)

func SolutionA() string {
	f, err := os.Open("./day_five.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stacks, scanner := GetStacks(f)

	for scanner.Scan() {
		instr := parseInstruction(scanner.Text())
		doInstruction(stacks, instr)
	}

	var sb strings.Builder
	for _, s := range stacks {
		if len(s) == 0 {
			continue
		}
		sb.WriteRune(s[len(s)-1])
	}

	return sb.String()
}

func SolutionB() string {
	f, err := os.Open("./day_five.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stacks, scanner := GetStacks(f)

	for scanner.Scan() {
		instr := parseInstruction(scanner.Text())
		doMultiInstruction(stacks, instr)
	}

	var sb strings.Builder
	for _, s := range stacks {
		if len(s) == 0 {
			continue
		}
		sb.WriteRune(s[len(s)-1])
	}

	return sb.String()
}

func GetStacks(f *os.File) (stacks, *bufio.Scanner) {
	stackInput, scanner := getStackInput(f)
	numberOfStacks := getNCols(stackInput[len(stackInput)-1])
	stacks := make([]stack, numberOfStacks+stackOffset)

	for _, line := range stackInput[:len(stackInput)-1] {
		for i := 0; i < numberOfStacks; i++ {
			if line[i*colWidth+runeOffset] != ' ' {
				stacks[i+stackOffset] = append(stacks[i+stackOffset], rune(line[i*colWidth+runeOffset]))
			}
		}
	}

	for i := range stacks {
		stacks[i].Reverse()
	}
	return stacks, scanner
}

func getStackInput(f io.Reader) ([]string, *bufio.Scanner) {
	scanner := bufio.NewScanner(f)
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == utils.EmptyString {
			break
		}
		input = append(input, line)
	}
	return input, scanner
}

func getNCols(s string) int {
	s = strings.TrimSpace(s)
	last := s[len(s)-1]
	i, err := strconv.Atoi(string(last))
	if err != nil {
		panic(err)
	}
	return i
}

type stacks []stack
type stack []rune

func (s stack) String() string {
	return string(s)
}

func (s stack) Reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (s stack) pop() (stack, rune) {
	popped := s[len(s)-1]
	return s[:len(s)-1], popped
}

func (s stack) push(r rune) stack {
	return append(s, r)
}

func (s stack) popMulti(i int) (stack, stack) {
	return s[:len(s)-i], s[len(s)-i:]
}

func (s stack) pushMulti(n stack) stack {
	return append(s, n...)
}

type instruction struct {
	quantity, from, to int
}

func parseInstruction(s string) instruction {
	split := strings.Split(s, " ")
	return instruction{instructionToInt(split[1]), instructionToInt(split[3]), instructionToInt(split[5])}
}

func instructionToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func doInstruction(stacks stacks, instr instruction) {
	for i := 0; i < instr.quantity; i++ {
		poppedStack, popped := stacks[instr.from].pop()
		stacks[instr.from] = poppedStack
		stacks[instr.to] = stacks[instr.to].push(popped)
	}
}

func doMultiInstruction(stacks stacks, instr instruction) {
	poppedStack, popped := stacks[instr.from].popMulti(instr.quantity)
	stacks[instr.from] = poppedStack
	stacks[instr.to] = stacks[instr.to].pushMulti(popped)
}
