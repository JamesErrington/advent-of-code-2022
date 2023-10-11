package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Option int

const (
	ROCK Option = iota
	PAPER
	SCISSORS
)

type Outcome int

const (
	WIN Outcome = iota
	DRAW
	LOSE
)

const (
	WIN_POINTS      = 6
	DRAW_POINTS     = 3
	ROCK_POINTS     = 1
	PAPER_POINTS    = 2
	SCISSORS_POINTS = 3
)

func main() {
	path, _ := filepath.Abs("../input/02/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		opponent_option, outcome := option_from_letter(line[0]), outcome_from_letter(line[1])

		switch outcome {
		case WIN:
			score += WIN_POINTS
			switch opponent_option {
			case ROCK:
				score += PAPER_POINTS
			case PAPER:
				score += SCISSORS_POINTS
			case SCISSORS:
				score += ROCK_POINTS
			}
		case DRAW:
			score += DRAW_POINTS
			switch opponent_option {
			case ROCK:
				score += ROCK_POINTS
			case PAPER:
				score += PAPER_POINTS
			case SCISSORS:
				score += SCISSORS_POINTS
			}
		case LOSE:
			switch opponent_option {
			case ROCK:
				score += SCISSORS_POINTS
			case PAPER:
				score += ROCK_POINTS
			case SCISSORS:
				score += PAPER_POINTS
			}
		}
	}
	fmt.Println(score)
}

func option_from_letter(letter string) Option {
	switch letter {
	case "A":
		return ROCK
	case "B":
		return PAPER
	default:
		return SCISSORS
	}
}

func outcome_from_letter(letter string) Outcome {
	switch letter {
	case "X":
		return LOSE
	case "Y":
		return DRAW
	default:
		return WIN
	}
}
