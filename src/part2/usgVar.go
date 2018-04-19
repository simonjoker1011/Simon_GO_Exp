package main

import (
	"fmt"
)

const myName = "simon"

func main() {
	const myNum = 7

	fmt.Printf("My name: %v\n", myName)

	//	error: const is final
	//	myName = "joker"
	//	fmt.Printf("My name: %v\n", myName)

	anotherF()

	var i int
	i = 0
	fmt.Printf("i: %v\n", i)

	j, k := 3, 4
	fmt.Printf("j: %v, k: %v\n", j, k)

	var l = "str var"

	fmt.Printf("l: %v", l)
}

func anotherF() {
	fmt.Printf("my name: %v\n", myName)
	//	fmt.Printf("my num: %v\n", myNum)
}
