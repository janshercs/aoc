package daythirteen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	in := `[1]
[2]

[3]
[4]
`

	assert.Equal(t, []pair{
		{[]interface{}{float64(1)}, []interface{}{float64(2)}},
		{[]interface{}{float64(3)}, []interface{}{float64(4)}},
	}, parseInput(strings.NewReader(in)))

}

func TestParseLine(t *testing.T) {
	assert.Equal(t, []interface{}{}, parseLine("[]"))
	assert.Equal(t, []interface{}{float64(1)}, parseLine("[1]"))
	assert.Equal(t, []interface{}{float64(1), float64(32), []interface{}{float64(3)}}, parseLine("[1, 32, [3]]"))
}
