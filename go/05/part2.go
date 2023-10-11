package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Instruction struct {
	Value int
	From  int
	To    int
}

type Stack []string

func main() {
	path, _ := filepath.Abs("../input/05/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	var stacks map[int]Stack
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			stacks = parse_stacks(lines)
			lines = nil
			continue
		}
		lines = append(lines, scanner.Text())
	}
	instructions := parse_instructions(lines)

	for _, instruction := range instructions {
		elements := stacks[instruction.From][len(stacks[instruction.From])-instruction.Value:]
		stacks[instruction.From] = stacks[instruction.From][:len(stacks[instruction.From])-instruction.Value]
		stacks[instruction.To] = append(stacks[instruction.To], elements...)

	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Print("\n")
}

func parse_stacks(lines []string) map[int]Stack {
	var instructions []Instruction

	for _, line := range lines[:len(lines)-1] {
		stack_number := 1
		for j, char := range line {
			if j == 3 || j > 4 && j%4 == 0 {
				stack_number++
			}

			if char >= 'A' && char <= 'Z' {
				instructions = append(instructions, Instruction{int(char), -1, stack_number})
			}
		}
	}

	stacks := make(map[int]Stack)

	for i := len(instructions) - 1; i >= 0; i-- {
		instruction := instructions[i]
		target := instruction.To

		if stacks[target] == nil {
			stacks[target] = make(Stack, 0)
		}

		stacks[target] = append(stacks[target], string(instruction.Value))
	}

	return stacks
}

func parse_instructions(lines []string) []Instruction {
	var instructions []Instruction
	for _, line := range lines {
		parts := strings.Split(line, " ")
		value, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		instructions = append(instructions, Instruction{value, from, to})
	}
	return instructions
}
