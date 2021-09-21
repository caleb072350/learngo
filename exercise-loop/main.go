package main

import (
	"exercise_loop/loop_over_slice"
	"exercise_loop/loop_rand"
	"fmt"
)

func main() {
	var sum int

	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	i := 1

	for {
		if i > 5 {
			break
		}

		sum += i

		fmt.Println(i, "---->", sum)

		i++
	}
	fmt.Println(sum)

	// use lib package
	loop_over_slice.Test()

	loop_rand.TestRand()
}
