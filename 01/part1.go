package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./01/question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	best_elf := 0
	current_elf := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if current_elf > best_elf {
				best_elf = current_elf
			}
			current_elf = 0
			continue
		}

		calories, _ := strconv.Atoi(line)
		current_elf += calories
	}

	if current_elf > best_elf {
		best_elf = current_elf
	}

	fmt.Println(best_elf)
}
