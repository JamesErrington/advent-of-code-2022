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
		opponent_option, my_play := option_from_letter(line[0]), option_from_letter(line[1])

		switch my_play {
		case ROCK:
			score += ROCK_POINTS
			switch opponent_option {
			case ROCK:
				score += DRAW_POINTS
			case SCISSORS:
				score += WIN_POINTS
			}
		case PAPER:
			score += PAPER_POINTS
			switch opponent_option {
			case ROCK:
				score += WIN_POINTS
			case PAPER:
				score += DRAW_POINTS
			}
		case SCISSORS:
			score += SCISSORS_POINTS
			switch opponent_option {
			case PAPER:
				score += WIN_POINTS
			case SCISSORS:
				score += DRAW_POINTS
			}
		}
	}
	fmt.Println(score)
}

func option_from_letter(letter string) Option {
	switch letter {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	default:
		return SCISSORS
	}
}
