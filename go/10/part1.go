package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path, _ := filepath.Abs("../input/10/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	X := 1
	cycle := 1
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		switch line[0] {
		case "noop":
			cycle++
			sum += check_cycle(cycle, X)
		case "addx":
			cycle++
			sum += check_cycle(cycle, X)
			cycle++
			value, _ := strconv.Atoi(line[1])
			X += value
			sum += check_cycle(cycle, X)
		}
	}

	fmt.Println(sum)
}

func check_cycle(cycle int, X int) int {
	if cycle == 20 || (cycle-20)%40 == 0 {
		return cycle * X
	}

	return 0
}
