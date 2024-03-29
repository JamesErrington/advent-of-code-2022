package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("../input/03/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count := make(map[rune]int)
		for _, item := range scanner.Text() {
			count[item] = 1
		}
		scanner.Scan()

		for _, item := range scanner.Text() {
			if count[item] == 1 {
				count[item] += 1
			}
		}
		scanner.Scan()

		for _, item := range scanner.Text() {
			if count[item] == 2 {
				score += item_priority(item)
				break
			}
		}
	}

	fmt.Println(score)
}

func item_priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 96
	}

	return int(item) - 38
}
