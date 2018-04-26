package main

import (
	"fmt"
)

func switchFunc(a int, b string) (string, int, bool) {
	fmt.Printf("multi param/return function\n")
	fmt.Printf("first param: %v, second param: %v \n", a, b)
	return "1st rtn", 2, true
}
