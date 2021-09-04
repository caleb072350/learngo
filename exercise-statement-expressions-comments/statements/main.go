package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println("cpus")
	fmt.Println(("the machine"))

	if 5 > 1 {
		fmt.Println("bigger")
	}

	if 2 > 3 {
		fmt.Println("smaller")
	}
}
