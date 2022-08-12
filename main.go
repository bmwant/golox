package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		if line == "" {
			fmt.Println("Goodbye")
			break
		}
		run(line)
	}
}

func error_(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Fprintln(os.Stderr, line, where, message)
	hadError := true
	_ = hadError
}

func runFile(s string) {
	fmt.Print(s)
}

func run(line string) {
	fmt.Println(line)
}

func main() {
	args := os.Args
	if len(args) > 1 {
		filename := args[1]
		data, err := os.ReadFile(filename)
		check(err)
		runFile(string(data))
	} else {
		runPrompt()
	}

	error_(0, "Just a test error message")
}
