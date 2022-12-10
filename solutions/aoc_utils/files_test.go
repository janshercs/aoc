package utils_test

import (
	utils "solutions/aoc_utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAndReadFile(t *testing.T) {
	f, err := utils.OpenFile("./test_file.txt")
	assert.NoError(t, err)
	defer f.Close()

	l, err := utils.ReadFile(f)
	assert.NoError(t, err)
	assert.Equal(t, `hi
bye
`, string(l))
}

func TestReadLine(t *testing.T) {
	f, err := utils.OpenFile("./test_file.txt")
	assert.NoError(t, err)
	defer f.Close()

	scanner := utils.ReadLine(f)

	scanner.Scan()
	line := scanner.Text()
	assert.Equal(t, "hi", line)

	scanner.Scan()
	line = scanner.Text()
	assert.Equal(t, "bye", line)

	scanner.Scan()
	line = scanner.Text()
	assert.Equal(t, "", line)
}
