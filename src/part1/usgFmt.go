package main

import (
	"fmt"
)

func main() {
	fmt.Printf("ABC")
	//	new line
	fmt.Println("-----new line-----")
	// \n as well
	fmt.Print("DEF\n")
	//concat
	fmt.Print("123" + "456" + "\n")

	//	format:

	//	string
	fmt.Printf("substution %s, %s", "here", "there\n")
	//	boolean
	fmt.Printf("boolean: %t\n", true)
	//	int, base10
	fmt.Printf("two: %d, five: %d\n", 2, 5)
	//	universal
	fmt.Printf("string: %v, boolean: %v, int: %v, double: %v", "str", false, 1, 3.14)
	//	format error
	fmt.Printf("\ntwo: %t", "string")
}
