package utils_test

import (
	utils "solutions/aoc_utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphabetPosition(t *testing.T) {
	i, err := utils.AlphabetPosition('a')
	assert.NoError(t, err)
	assert.Equal(t, 1, i)

	i, err = utils.AlphabetPosition('A')
	assert.NoError(t, err)
	assert.Equal(t, 1, i)

	i, err = utils.AlphabetPosition('p')
	assert.NoError(t, err)
	assert.Equal(t, 16, i)
}

func TestGetAlphabetInPosition(t *testing.T) {
	A, err := utils.GetAlphabetInPosition(1)
	assert.NoError(t, err)
	assert.Equal(t, 'A', A)
	Z, _ := utils.GetAlphabetInPosition(26)
	assert.Equal(t, 'Z', Z)
}
