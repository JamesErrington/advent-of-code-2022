package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	UP    = "U"
	RIGHT = "R"
	DOWN  = "D"
	LEFT  = "L"
)

func main() {
	file, err := os.Open("./09/part1question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	run_simulation(scanner)
}

func run_simulation(source *bufio.Scanner) {
	head_row, head_col := 0, 0
	tail_row, tail_col := 0, 0
	visited := map[string]int{"0,0": -1}

	for source.Scan() {
		line := strings.Split(source.Text(), " ")

		step_size, _ := strconv.Atoi(line[1])
		for i := 0; i < step_size; i++ {
			switch line[0] {
			case UP:
				head_row += 1
			case RIGHT:
				head_col += 1
			case DOWN:
				head_row -= 1
			case LEFT:
				head_col -= 1
			default:
				panic("Unknown direction: " + line[0])
			}

			v_diff, h_diff := head_row-tail_row, head_col-tail_col

			if math.Abs(float64(v_diff)) <= 1 && math.Abs(float64(h_diff)) <= 1 {
				continue
			}

			switch {
			case (v_diff > 1 && h_diff > 0) || (h_diff > 1 && v_diff > 0):
				tail_col += 1
				tail_row += 1
			case (v_diff > 1 && h_diff < 0) || (h_diff < -1 && v_diff > 0):
				tail_col -= 1
				tail_row += 1
			case (h_diff < -1 && v_diff < 0) || (v_diff < -1 && h_diff < 0):
				tail_col -= 1
				tail_row -= 1
			case (h_diff > 1 && v_diff < 0) || (v_diff < -1 && h_diff > 0):
				tail_col += 1
				tail_row -= 1
			case h_diff > 1:
				tail_col += 1
			case h_diff < -1:
				tail_col -= 1
			case v_diff > 1:
				tail_row += 1
			case v_diff < -1:
				tail_row -= 1
			}

			visited[fmt.Sprintf("%d,%d", tail_col, tail_row)] = i
		}
	}

	fmt.Println(len(visited))
}
