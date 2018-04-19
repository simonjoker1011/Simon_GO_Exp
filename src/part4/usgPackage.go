package main

import (
	"extras"
	"extras/extra1"
	"fmt"
)

func main() {
	//	for other package func, must capt at first char
	fmt.Println(extras.GetPI())

	//	this will fail
	//	fmt.Println(extras.getPI())

	fmt.Print(extra1.GetExp())

}
