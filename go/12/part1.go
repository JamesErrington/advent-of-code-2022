package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("../input/12/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	terrain := [][]int{}
	var N_n, S_x, S_y, E_x, E_y int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, len(line))
		for i, char := range line {
			N_n++
			row[i] = int(char)
			if char == 'S' {
				S_x = i
				S_y = len(terrain)
				row[i] = int('a')
			}

			if char == 'E' {
				E_x = i
				E_y = len(terrain)
				row[i] = int('z')
			}
		}
		terrain = append(terrain, row)
	}
	width := len(terrain[0])
	S_i := index(S_x, S_y, width)
	E_i := index(E_x, E_y, width)

	nodes := make([]int, 0, N_n)
	edges := make(map[int][]int)
	for y, row := range terrain {
		for x, item := range row {
			nodes = append(nodes, item)

			e := make([]int, 0)
			if y > 0 && terrain[y-1][x] <= item+1 {
				e = append(e, index(x, y-1, width))
			}

			if y < len(terrain)-1 && terrain[y+1][x] <= item+1 {
				e = append(e, index(x, y+1, width))
			}

			if x > 0 && terrain[y][x-1] <= item+1 {
				e = append(e, index(x-1, y, width))
			}

			if x < width-1 && terrain[y][x+1] <= item+1 {
				e = append(e, index(x+1, y, width))
			}

			edges[index(x, y, width)] = e
		}
	}

	openSet := []int{S_i}

	cameFrom := make(map[int]int)

	gScore := make(map[int]int, N_n)
	for i, _ := range nodes {
		gScore[i] = math.MaxInt
	}
	gScore[S_i] = 0

	fScore := make(map[int]float64, N_n)
	for i, _ := range nodes {
		fScore[i] = math.MaxInt
	}
	fScore[S_i] = h(S_i, E_i, width)

	for len(openSet) > 0 {
		current, c_i := openSet[0], 0
		for i, node := range openSet {
			if fScore[node] < fScore[current] {
				current = node
				c_i = i
			}
		}

		if current == E_i {
			total_path := []int{current}
			for {
				value, ok := cameFrom[current]
				if !ok {
					break
				}

				current = value
				total_path = append(total_path, current)
			}

			fmt.Println(len(total_path) - 1)
			break
		}

		openSet = append(openSet[:c_i], openSet[c_i+1:]...)
		for _, neighbour := range edges[current] {
			tentative_gScore := gScore[current] + 1
			if tentative_gScore < gScore[neighbour] {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentative_gScore
				fScore[neighbour] = float64(tentative_gScore) + h(current, neighbour, width)

				add := true
				for _, node := range openSet {
					if node == neighbour {
						add = false
						break
					}
				}
				if add {
					openSet = append(openSet, neighbour)
				}
			}
		}
	}
}

func index(x int, y int, w int) int {
	return y*w + x
}

func h(i int, e int, w int) float64 {
	S_x, S_y := float64(i%w), float64(i/w)
	E_x, E_y := float64(e%w), float64(e/w)

	return math.Sqrt(math.Pow(S_x-E_x, 2) + math.Pow(S_y-E_y, 2))
}
