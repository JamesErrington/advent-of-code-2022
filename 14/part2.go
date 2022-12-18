package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	start_x, start_y, end_x, end_y int
}

type Row []string

type Grid []Row

var (
	SAND_SOURCE_X = 500
)

const (
	SAND_SOURCE_Y = 0
	AIR           = "."
	ROCK          = "#"
	SAND_SOURCE   = "+"
	SAND          = "o"
)

func main() {
	file, err := os.Open("./14/question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []Line
	min_x, min_y := SAND_SOURCE_X, SAND_SOURCE_Y
	max_x, max_y := SAND_SOURCE_X, SAND_SOURCE_Y

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")

		for i := 0; i < len(parts)-1; i += 1 {
			start, end := strings.Split(parts[i], ","), strings.Split(parts[i+1], ",")
			start_x, _ := strconv.Atoi(start[0])
			start_y, _ := strconv.Atoi(start[1])
			end_x, _ := strconv.Atoi(end[0])
			end_y, _ := strconv.Atoi(end[1])

			if start_x < min_x || end_x < min_x {
				min_x = min(start_x, end_x)
			}

			if start_x > max_x || end_x > max_x {
				max_x = max(start_x, end_x)
			}

			if start_y < min_y || end_y < min_y {
				min_y = min(start_y, end_y)
			}

			if start_y > max_y || end_y > max_y {
				max_y = max(start_y, end_y)
			}

			lines = append(lines, Line{start_x, start_y, end_x, end_y})
		}
	}

	grid := make_grid(lines, min_x, max_x, min_y, max_y)

	i := 0
	continue_loop := true
	for ; continue_loop; i += 1 {
		sand_x, sand_y := SAND_SOURCE_X-min_x, SAND_SOURCE_Y-min_y

		for {
			replace_self := grid[sand_y][sand_x] != SAND_SOURCE

			below := grid[sand_y+1][sand_x]
			if below == ROCK || below == SAND {
				if sand_x == 0 {
					grid = extend_left(grid)
					sand_x += 1
					min_x -= 1
				}
				below_left := grid[sand_y+1][sand_x-1]
				if below_left == AIR {
					if replace_self {
						grid[sand_y][sand_x] = AIR
					}
					sand_y += 1
					sand_x -= 1
					grid[sand_y][sand_x] = SAND
					continue
				}

				if sand_x == len(grid[0])-1 {
					grid = extend_right(grid)
				}
				below_right := grid[sand_y+1][sand_x+1]
				if below_right == AIR {
					if replace_self {
						grid[sand_y][sand_x] = AIR
					}
					sand_y += 1
					sand_x += 1
					grid[sand_y][sand_x] = SAND
					continue
				}

				if sand_x == SAND_SOURCE_X-min_x && sand_y == SAND_SOURCE_Y-min_y {
					continue_loop = false
				}

				break
			}

			if replace_self {
				grid[sand_y][sand_x] = AIR
			}
			sand_y += 1
			grid[sand_y][sand_x] = SAND
		}
	}

	fmt.Println(i)
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func is_rock(lines []Line, x int, y int) bool {
	for _, line := range lines {
		min_x := min(line.start_x, line.end_x)
		max_x := max(line.start_x, line.end_x)
		min_y := min(line.start_y, line.end_y)
		max_y := max(line.start_y, line.end_y)
		if x >= min_x && x <= max_x && y >= min_y && y <= max_y {
			return true
		}
	}
	return false
}

func make_grid(lines []Line, min_x int, max_x int, min_y int, max_y int) Grid {
	grid := make(Grid, 0, max_y-min_y+1+2)
	for y := min_y; y <= max_y; y += 1 {
		row := make(Row, 0, max_x-min_x+1)
		for x := min_x; x <= max_x; x += 1 {
			switch {
			case x == SAND_SOURCE_X && y == SAND_SOURCE_Y:
				row = append(row, SAND_SOURCE)
			case is_rock(lines, x, y):
				row = append(row, ROCK)
			default:
				row = append(row, AIR)
			}
		}
		grid = append(grid, row)
	}

	air_gap := make(Row, 0, max_x-min_x)
	floor := make(Row, 0, max_x-min_x)
	for x := 0; x < max_x-min_x+1; x += 1 {
		air_gap = append(air_gap, AIR)
		floor = append(floor, ROCK)
	}
	grid = append(grid, air_gap, floor)

	return grid
}

func print_grid(grid Grid) {
	for _, row := range grid {
		for _, square := range row {
			fmt.Print(square)
		}
		fmt.Println()
	}
	fmt.Println()
}

func extend_left(grid Grid) Grid {
	new_grid := make(Grid, 0, len(grid))
	for y, row := range grid {
		new_row := make(Row, 0, len(row)+1)
		if y == len(grid)-1 {
			new_row = append(Row{ROCK}, row...)
		} else {
			new_row = append(Row{AIR}, row...)
		}
		new_grid = append(new_grid, new_row)
	}

	return new_grid
}

func extend_right(grid Grid) Grid {
	new_grid := make(Grid, 0, len(grid))
	for y, row := range grid {
		new_row := make(Row, 0, len(row)+1)
		if y == len(grid)-1 {
			new_row = append(row, ROCK)
		} else {
			new_row = append(row, AIR)
		}
		new_grid = append(new_grid, new_row)
	}

	return new_grid
}
