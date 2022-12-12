package dayfour

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getAssignmentPairs(s string) []assignment {
	var assignments []assignment
	elves := strings.Split(s, ",")
	for _, elf := range elves {
		assignments = append(assignments, getAssignment(elf))
	}

	return assignments
}

type assignment struct {
	start int
	end   int
}

func getAssignment(s string) assignment {
	startEnd := strings.Split(s, "-")
	start, err := strconv.Atoi(startEnd[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(startEnd[1])
	if err != nil {
		panic(err)
	}
	return assignment{start: start, end: end}
}

// returns assignment with smaller start, if both assignment has same start, it will return assignment with the bigger end
func sortByStart(assignments []assignment) (assignment, assignment) {
	if assignments[1].start < assignments[0].start {
		return assignments[1], assignments[0]
	}

	if assignments[1].start == assignments[0].start {
		return sortByEnd(assignments)
	}

	return assignments[0], assignments[1]
}

func sortByEnd(assignments []assignment) (assignment, assignment) {
	if assignments[0].end >= assignments[1].end {
		return assignments[0], assignments[1]
	}
	return assignments[1], assignments[0]
}

func isContaining(smaller, bigger assignment) bool {
	return smaller.end >= bigger.end
}

func isOverlapping(smaller, bigger assignment) bool {
	return smaller.end >= bigger.start
}

func SolutionA() int {
	f, err := os.Open("./day_four.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		if isContaining(sortByStart(getAssignmentPairs(scanner.Text()))) {
			count++
		}
	}
	return count
}

func SolutionB() int {
	f, err := os.Open("./day_four.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		if isOverlapping(sortByStart(getAssignmentPairs(scanner.Text()))) {
			count++
		}
	}
	return count
}
