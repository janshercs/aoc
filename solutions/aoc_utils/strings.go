package utils

import (
	"strconv"
	"strings"
)

func GetLastIntInString(s string) int {
	input := strings.Split(s, " ")
	divisor := input[len(input)-1]

	i, err := strconv.Atoi(divisor)
	if err != nil {
		panic(err)
	}
	return i
}
