package day02

import (
	"fmt"
	"strings"
)

type move int
type strategy int

const (
	ROCK move = 1 + iota
	PAPER
	SCISSORS
)

const (
	LOSE strategy = iota
	DRAW
	WIN
)

type responseMove struct {
	winningRespone move
	losingResponse move
}

var inputToEnemyMove = map[string]move{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var inputToOwnMove = map[string]move{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var inputToOwnStrategy = map[string]strategy{
	"X": LOSE,
	"Y": DRAW,
	"Z": WIN,
}

var enemyMoveToResponses = map[move]responseMove{
	ROCK:     {PAPER, SCISSORS},
	PAPER:    {SCISSORS, ROCK},
	SCISSORS: {ROCK, PAPER},
}

func Solve(rawInput []string) (string, string) {
	firstSolutionTally := 0
	secondSolutionTally := 0
	for _, v := range rawInput {
		splitInput := strings.Fields(v)
		enemyMove := inputToEnemyMove[splitInput[0]]
		ownMovePart1 := inputToOwnMove[splitInput[1]]

		ownStrategyPart2 := inputToOwnStrategy[splitInput[1]]
		ownMovePart2 := determineMoveByStrategy(enemyMove, ownStrategyPart2)

		firstSolutionTally += calculateScoreIncrease(enemyMove, ownMovePart1)
		secondSolutionTally += calculateScoreIncrease(enemyMove, ownMovePart2)
	}

	return fmt.Sprint(firstSolutionTally), fmt.Sprint(secondSolutionTally)
}

func calculateScoreIncrease(enemyMove move, ownMove move) int {
	tally := int(ownMove)
	if enemyMove == ownMove {
		tally += 3
	} else if enemyMoveToResponses[enemyMove].winningRespone == ownMove {
		tally += 6
	}
	return tally
}

func determineMoveByStrategy(enemyMove move, strategy strategy) move {
	responses := enemyMoveToResponses[enemyMove]
	if strategy == LOSE {
		return responses.losingResponse
	} else if strategy == WIN {
		return responses.winningRespone
	}
	return enemyMove
}
