package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("../input/06/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := scanner.Text()
	buffer := NewBuffer[rune](14)
	for i, char := range line {
		buffer.Push(char)

		if i >= 14 {
			duplicate := false
			for j, x := range buffer.Data {
				for k, y := range buffer.Data {
					if j != k && x == y {
						duplicate = true
					}
				}
			}

			if !duplicate {
				fmt.Println(i + 1)
				break
			}
		}
	}
}

type Buffer[T interface{}] struct {
	Data  []T
	index int
}

func NewBuffer[T interface{}](capacity int) Buffer[T] {
	return Buffer[T]{make([]T, capacity, capacity), 0}
}

func (buffer *Buffer[T]) Push(value T) {
	buffer.Data[buffer.index] = value
	buffer.index++
	buffer.index %= cap(buffer.Data)
}
