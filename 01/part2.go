package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./01/part1question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	elves := &IntHeap{}
	current_elf := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			heap.Push(elves, -current_elf)
			current_elf = 0
			continue
		}

		calories, _ := strconv.Atoi(line)
		current_elf += calories
	}

	heap.Push(elves, -current_elf)

	total := -heap.Pop(elves).(int)
	for i := 0; i < 2; i++ {
		total += -heap.Pop(elves).(int)
	}
	fmt.Println(total)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
