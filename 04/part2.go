package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./04/question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		left, right := strings.Split(line[0], "-"), strings.Split(line[1], "-")

		ll, lu, rl, ru := parse_bounds(left, right)
		if (ll <= rl && lu >= rl) || (rl <= ll && ru >= ll) {
			count += 1
		}
	}

	fmt.Println(count)
}

func parse_bounds(left []string, right []string) (int, int, int, int) {
	ll, _ := strconv.Atoi(left[0])
	lu, _ := strconv.Atoi(left[1])
	rl, _ := strconv.Atoi(right[0])
	ru, _ := strconv.Atoi(right[1])

	return ll, lu, rl, ru
}
