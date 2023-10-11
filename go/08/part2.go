package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	path, _ := filepath.Abs("../input/08/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, len(line))
		for i, char := range line {
			cell, _ := strconv.Atoi(string(char))
			row[i] = cell
		}

		grid = append(grid, row)
	}

	best := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			height := grid[i][j]

			count := check_left(grid, height, i, j-1) * check_right(grid, height, i, j+1) * check_up(grid, height, i-1, j) * check_down(grid, height, i+1, j)
			if count > best {
				best = count
			}
		}
	}

	fmt.Println(best)
}

func check_left(grid [][]int, height int, i int, j int) int {
	if j == -1 {
		return 0
	}

	if grid[i][j] < height {
		return 1 + check_left(grid, height, i, j-1)
	}

	return 1
}

func check_right(grid [][]int, height int, i int, j int) int {
	if j == len(grid[0]) {
		return 0
	}

	if grid[i][j] < height {
		return 1 + check_right(grid, height, i, j+1)
	}

	return 1
}

func check_up(grid [][]int, height int, i int, j int) int {
	if i == -1 {
		return 0
	}

	if grid[i][j] < height {
		return 1 + check_up(grid, height, i-1, j)
	}

	return 1
}

func check_down(grid [][]int, height int, i int, j int) int {
	if i == len(grid) {
		return 0
	}

	if grid[i][j] < height {
		return 1 + check_down(grid, height, i+1, j)
	}

	return 1
}
