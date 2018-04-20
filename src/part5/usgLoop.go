package main

import (
	"fmt"
)

func main() {
	a, b, c := switchFunc(0, "hi")
	fmt.Printf("1st rtn: %v, 2ed rtn: %v, 3rd rtn: %v\n", a, b, c)

	for i := 1; i < 10; i++ {
		fmt.Printf("step: %v\n", i)
	}
	// do while
	j := 0
	for {
		fmt.Printf("j: %v\n", j)
		if j > 10 {
			break
		}
		j++
	}
	// while
	k := 0
	for k < 10 {
		fmt.Printf("k: %v\n", k)
		k++
	}

	//	infinity loop
	//
	//	for {
	//	}
}
