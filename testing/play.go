package main

import "fmt"

type TokenType int

const (
	ONE TokenType = iota
	TWO
)

type Token struct {
	Name      string
	TokenType TokenType
}

func main() {
	a := TWO
	b := Token{"Name", TWO}
	b1 := Token{"B1", ONE}
	fmt.Println("This is my testing", a, b)
	fmt.Println(b1, b1.TokenType)
	fmt.Printf("%T %T\n", b.TokenType, b1.TokenType)
}
