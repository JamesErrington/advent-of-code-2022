package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./08/question.txt")
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

	count := len(grid)*4 - 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			height := grid[i][j]

			visible := check_left(grid, height, i, j-1) || check_right(grid, height, i, j+1) || check_up(grid, height, i-1, j) || check_down(grid, height, i+1, j)
			if visible {
				count++
			}
		}
	}

	fmt.Println(count)
}

func check_left(grid [][]int, height int, i int, j int) bool {
	if j == -1 {
		return true
	}

	if grid[i][j] < height {
		return check_left(grid, height, i, j-1)
	}

	return false
}

func check_right(grid [][]int, height int, i int, j int) bool {
	if j == len(grid[0]) {
		return true
	}

	if grid[i][j] < height {
		return check_right(grid, height, i, j+1)
	}

	return false
}

func check_up(grid [][]int, height int, i int, j int) bool {
	if i == -1 {
		return true
	}

	if grid[i][j] < height {
		return check_up(grid, height, i-1, j)
	}

	return false
}

func check_down(grid [][]int, height int, i int, j int) bool {
	if i == len(grid) {
		return true
	}

	if grid[i][j] < height {
		return check_down(grid, height, i+1, j)
	}

	return false
}
