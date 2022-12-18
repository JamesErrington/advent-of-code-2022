package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("ERROR! No day argument given. Usage: make [day]")
		os.Exit(1)
	}
	day := args[1]

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("%s/%s", wd, day)

	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	os.Create(fmt.Sprintf("%s/%s", path, "example.txt"))
	os.Create(fmt.Sprintf("%s/%s", path, "question.txt"))
	file, err := os.Create(fmt.Sprintf("%s/%s", path, "part1.go"))
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("package main\n\nfunc main() {\nfile, err := os.Open(\"./%s/example.txt\")\nif err != nil {\npanic(err)\n}\ndefer file.Close()\n\nscanner := bufio.NewScanner(file)\n}\n", day))
	writer.Flush()
}
