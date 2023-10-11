package main

import (
	"bufio"
	"fmt"
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
	var N_n, E_x, E_y int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, len(line))
		for i, char := range line {
			N_n++
			row[i] = int(char)
			if char == 'S' {
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
	E_i := index(E_x, E_y, width)

	nodes := make([]int, 0, N_n)
	edges := make(map[int][]int)
	as := make([]int, 0)
	for y, row := range terrain {
		for x, item := range row {
			nodes = append(nodes, item)

			e := make([]int, 0)
			if y > 0 && terrain[y-1][x] >= item-1 {
				e = append(e, index(x, y-1, width))
			}

			if y < len(terrain)-1 && terrain[y+1][x] >= item-1 {
				e = append(e, index(x, y+1, width))
			}

			if x > 0 && terrain[y][x-1] >= item-1 {
				e = append(e, index(x-1, y, width))
			}

			if x < width-1 && terrain[y][x+1] >= item-1 {
				e = append(e, index(x+1, y, width))
			}

			i := index(x, y, width)
			edges[i] = e
			if item == 'a' {
				as = append(as, i)
			}
		}
	}

	dist := map[int]int{}
	Q := make([]int, N_n)

	for v, _ := range nodes {
		dist[v] = 99999
		Q[v] = v
	}
	dist[E_i] = 0

	for len(Q) > 0 {
		u, u_i := Q[0], 0
		for i, node := range Q {
			if dist[node] < dist[u] {
				u = node
				u_i = i
			}
		}

		Q = append(Q[:u_i], Q[u_i+1:]...)
		for _, v := range edges[u] {
			found := false
			for _, q := range Q {
				if q == v {
					found = true
				}
			}
			if !found {
				continue
			}

			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
			}
		}
	}

	min := dist[0]
	for i, h := range nodes {
		if h == 'a' {
			if dist[i] < min {
				min = dist[i]
			}
		}
	}
	fmt.Println(min)
}

func index(x int, y int, w int) int {
	return y*w + x
}
