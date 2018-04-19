package main

import (
	"fmt"
	"runtime"
)

var con = float64(10)

func main() {
	var arr = []string{"a", "b", "c", "d"}

	fmt.Println(arr)

	for i := range arr {
		fmt.Println(i, arr[i])
	}

	fmt.Println(runtime.GOOS, runtime.Version())
}
