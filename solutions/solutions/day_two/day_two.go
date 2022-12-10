package daytwo

import (
	"fmt"
	utils "solutions/aoc_utils"
	"strings"
)

type outcome string
type shape interface {
	beats(string) bool
	String() string
	losesTo() shape
	winsOver() shape
}

type rockShape struct{}

func (r rockShape) beats(s string) bool {
	return s == "scissor"
}

func (r rockShape) winsOver() shape {
	return scissorShape{}
}

func (r rockShape) losesTo() shape {
	return paperShape{}
}

func (r rockShape) String() string {
	return "rock"
}

type paperShape struct{}

func (r paperShape) beats(s string) bool {
	return s == "rock"
}

func (r paperShape) winsOver() shape {
	return rockShape{}
}

func (r paperShape) losesTo() shape {
	return scissorShape{}
}

func (r paperShape) String() string {
	return "paper"
}

type scissorShape struct{}

func (r scissorShape) beats(s string) bool {
	return s == "paper"
}

func (r scissorShape) winsOver() shape {
	return paperShape{}
}

func (r scissorShape) losesTo() shape {
	return rockShape{}
}

func (r scissorShape) String() string {
	return "scissor"
}

const (
	win  outcome = "win"
	draw outcome = "draw"
	lose outcome = "lose"
)

var (
	rock     shape = rockShape{}
	paper    shape = paperShape{}
	scissors shape = scissorShape{}
)

var InputToShape = map[string]shape{
	"A": rock,
	"X": rock,
	"B": paper,
	"Y": paper,
	"C": scissors,
	"Z": scissors,
}

var InputToOutcome = map[string]outcome{
	"X": lose,
	"Y": draw,
	"Z": win,
}

var shapeScore = map[shape]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var outcomeScore = map[outcome]int{
	win:  6,
	draw: 3,
	lose: 0,
}

func Solution() (int, error) {
	f, err := utils.OpenFile("./day_2.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	defer f.Close()

	score := 0
	scanner := utils.ReadLine(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == utils.EmptyString {
			break
		}
		in := splitLine(line)
		score += MatchScore(in[0], in[1])
	}

	return score, nil
}

func SolutionB() (int, error) {
	f, err := utils.OpenFile("./day_2.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	defer f.Close()

	score := 0
	scanner := utils.ReadLine(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == utils.EmptyString {
			break
		}
		in := splitLine(line)
		score += MatchScoreFromOutcome(in[0], in[1])
	}

	return score, nil
}

func MatchScore(opp, self string) int {
	oc := EvaluateOutcome(InputToShape[opp], InputToShape[self])

	return outcomeScore[oc] + shapeScore[InputToShape[self]]
}

func EvaluateOutcome(opp, self shape) outcome {
	if opp == self {
		return draw
	}

	if self.beats(opp.String()) {
		return win
	}

	return lose
}

func splitLine(s string) []string {
	return strings.Split(s, " ")
}

func MatchScoreFromOutcome(opp, outcome string) int {
	return outcomeScore[InputToOutcome[outcome]] + shapeScore[EvaluateShape(InputToShape[opp], InputToOutcome[outcome])]
}

func EvaluateShape(opp shape, outcome outcome) shape {
	switch outcome {
	case draw:
		return opp
	case win:
		return opp.losesTo()
	default:
		return opp.winsOver()
	}
}
