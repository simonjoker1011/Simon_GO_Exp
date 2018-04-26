package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- switch usage ---")

	i := 1

	switch i {
	case 0, 1:
		fmt.Print("0 or 1\n")
	case 2:
		fmt.Print("2\n")
	}

	switch {
	case i > 0:
		fmt.Printf("i: (%v) >0\n", i)
	case i == 0:
		fmt.Printf("i: (%v) =0\n", i)
	case i < 0:
		fmt.Printf("i: (%v) <0\n", i)
	default:
		fmt.Printf("Default")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
}
