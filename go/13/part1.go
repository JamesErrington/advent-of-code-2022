package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type ElementType string

const (
	ELEMENT_LIST    ElementType = "LIST"
	ELEMENT_INTEGER ElementType = "INT"
)

type Element struct {
	_type   ElementType
	integer int
	list    []Element
}

func (e Element) is_int() bool {
	return e._type == ELEMENT_INTEGER
}

func (e Element) is_list() bool {
	return e._type == ELEMENT_LIST
}

func main() {
	path, _ := filepath.Abs("../input/13/question.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	index := 1
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left_str := scanner.Text()
		scanner.Scan()
		right_str := scanner.Text()
		scanner.Scan()

		left := parse_packet(left_str)
		right := parse_packet(right_str)

		order := compare(left, right)
		fmt.Println("Correct order:", order)
		if order == CORRECT_ORDER {
			sum += index
		}

		index += 1
	}

	fmt.Println(sum)
}

func parse_packet(input string) Element {
	element, _ := parse_element(input)
	return element
}

func parse_element(input string) (Element, int) {
	switch input[0] {
	case '[':
		list, length := parse_list(input)
		return Element{ELEMENT_LIST, len(list), list}, length
	default:
		integer, length := parse_integer(input)
		return Element{ELEMENT_INTEGER, integer, nil}, length
	}
}

func parse_list(input string) ([]Element, int) {
	list := make([]Element, 0)
	i := 1
	for i < len(input) {
		if input[i] == ']' {
			i += 1
			break
		} else if input[i] == ',' {
			i += 1
			continue
		}

		element, length := parse_element(input[i:])

		list = append(list, element)
		i += length
	}

	return list, i
}

func parse_integer(input string) (int, int) {
	length := 0
	for i, char := range input {
		if char < '0' || char > '9' {
			length = i
			break
		}
	}

	integer, _ := strconv.Atoi(input[:length])
	return integer, length
}

type SortResult int

const (
	INCORRECT_ORDER SortResult = iota
	CORRECT_ORDER
	EQUAL_ORDER
)

func compare(left Element, right Element) SortResult {
	switch {
	case left.is_int() && right.is_int():
		if left.integer < right.integer {
			return CORRECT_ORDER
		}

		if left.integer > right.integer {
			return INCORRECT_ORDER
		}

		return EQUAL_ORDER
	case left.is_list() && right.is_list():
		for i := 0; ; i += 1 {
			switch {
			case i == len(left.list) && i == len(right.list):
				return EQUAL_ORDER
			case i == len(left.list):
				return CORRECT_ORDER
			case i == len(right.list):
				return INCORRECT_ORDER
			}

			left_elem, right_elem := left.list[i], right.list[i]
			order := compare(left_elem, right_elem)
			if order != EQUAL_ORDER {
				return order
			}
		}
	case left.is_int():
		return compare(Element{ELEMENT_LIST, 1, []Element{left}}, right)
	case right.is_int():
		return compare(left, Element{ELEMENT_LIST, 1, []Element{right}})
	}

	panic("Should be unreachable?")
}
