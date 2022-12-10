package dayone_test

import (
	dayone "solutions/solutions/day_one"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	i, err := dayone.Solution()
	assert.NoError(t, err)
	assert.Equal(t, 203420, i)
}
