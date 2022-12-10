package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./03/question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		items_per_compartment := len(line) / 2
		left_items, right_items := line[:items_per_compartment], line[items_per_compartment:]

		duplicate := '0'
		for _, left_item := range left_items {
			for _, right_item := range right_items {
				if left_item == right_item {
					duplicate = right_item
				}
			}
		}

		score += item_priority(duplicate)
	}

	fmt.Println(score)
}

func item_priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 96
	}

	return int(item) - 38
}
