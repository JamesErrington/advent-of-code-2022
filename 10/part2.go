package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./10/question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cycle := 1
	X := 1
	display := [6][40]string{}
	draw_pixel(&display, cycle, X)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		switch line[0] {
		case "noop":
			cycle++
			draw_pixel(&display, cycle, X)
		case "addx":
			cycle++
			draw_pixel(&display, cycle, X)
			cycle++
			value, _ := strconv.Atoi(line[1])
			X += value
			draw_pixel(&display, cycle, X)
		}
	}

	for _, row := range display {
		fmt.Println(row)
	}
}

func draw_pixel(display *[6][40]string, cycle int, X int) {
	row := (cycle - 1) / 40
	column := (cycle - 1) % 40

	if row >= 6 {
		return
	}

	if X-1 == column || X == column || X+1 == column {
		display[row][column] = "#"
	} else {
		display[row][column] = "."
	}
}
