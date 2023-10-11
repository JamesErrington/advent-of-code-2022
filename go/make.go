package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	path := filepath.Join(wd, day)

	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	input_dir := filepath.Join(wd, "../input/", day)
	err = os.Mkdir(input_dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	os.Create(filepath.Join(input_dir, "example.txt"))
	os.Create(filepath.Join(input_dir, "question.txt"))
	file, err := os.Create(fmt.Sprintf("%s/%s", path, "part1.go"))
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("package main\n\nfunc main() {\npath, _ := filepath.Abs(\"../input/%s/question.txt\")\nfile, err := os.Open(path)\nif err != nil {\npanic(err)\n}\ndefer file.Close()\n\nscanner := bufio.NewScanner(file)\n}\n", day))
	writer.Flush()
}
