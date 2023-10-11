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
	path, _ := filepath.Abs("../input/04/question.txt")
	file, err := os.Open(path)
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
		if (ll <= rl && lu >= ru) || (rl <= ll && ru >= lu) {
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
