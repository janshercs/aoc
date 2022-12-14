package daysix

import (
	"bufio"
	"os"
)

func getMarkerPosition(s string, n int) int {
	tracker := map[byte]bool{}
	i, j := 0, 0

	for j-i < n {
		if !tracker[s[j]] {
			tracker[s[j]] = true
			j++
			continue
		}
		// while s[j] still in tracker, set char at i to be false, and increase
		for tracker[s[j]] {
			tracker[s[i]] = false
			i++
		}
		tracker[s[j]] = true
		j++
	}

	return j
}

func solutionA() int {
	f, err := os.Open("./day_six.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan() // only 1 line
	return getMarkerPosition(scanner.Text(), 4)
}

func solutionB() int {
	f, err := os.Open("./day_six.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan() // only 1 line
	return getMarkerPosition(scanner.Text(), 14)
}
