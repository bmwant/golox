package main

import "fmt"

type MyScan struct {
	Source string
	Start  int
}

func NewMyScan(source string) MyScan {
	return MyScan{
		Source: source,
		Start:  1,
	}
}

func (ms *MyScan) tryThis() bool {
	return ms.Start >= len(ms.Source)
}

func main_test() {
	fmt.Println("Here we go")
	ms := NewMyScan("This is the my scan")
	fmt.Println(ms, ms.tryThis())
	ms.Start = 500
	fmt.Println(ms, ms.tryThis())
	c := 'a'
	switch c {
	case 'b':
		fmt.Println("We are b")
	case 'c':
		fmt.Println("Is this possible")
	default:
		fmt.Println("Nothing matched")
	}
}
