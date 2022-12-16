package main

import (
	"bufio"
	"fmt"
	"os"
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

func (e Element) is_divider() bool {
	return e._type == ELEMENT_LIST &&
		len(e.list) == 1 &&
		e.list[0]._type == ELEMENT_LIST &&
		len(e.list[0].list) == 1 &&
		e.list[0].list[0]._type == ELEMENT_INTEGER &&
		(e.list[0].list[0].integer == 2 || e.list[0].list[0].integer == 6)
}

func main() {
	file, err := os.Open("./13/question.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list := []Element{parse_packet("[[2]]"), parse_packet("[[6]]")}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		packet_str := scanner.Text()
		if len(packet_str) == 0 {
			continue
		}

		packet := parse_packet(packet_str)

		for i, other := range list {
			order := compare(packet, other)

			if order == CORRECT_ORDER {
				list = append(list[:i+1], list[i:]...)
				list[i] = packet
				break
			}

			if order == INCORRECT_ORDER && i == len(list)-1 {
				list = append(list, packet)
			}
		}
	}

	product := 1
	for i, element := range list {
		if element.is_divider() {
			fmt.Println(element)
			product *= (i + 1)
		}
	}
	fmt.Println(product)
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
