package main

import "fmt"

func main() {
	var (
		myAge   byte
		herAge  byte
		ages    [2]byte
		tags    [5]string
		zero    [0]byte
		agesExp [2 + 1]byte
	)

	fmt.Printf("myAge :     %d\n", myAge)
	fmt.Printf("herAge:     %d\n", herAge)

	fmt.Printf("ages :      %#v\n", ages)
	fmt.Printf("tags    :   %#v\n", tags)
	fmt.Printf("zero    :   %#v\n", zero)
	fmt.Printf("agesExp :   %#v\n", agesExp)

	{
		var ages [2]int

		fmt.Printf("ages :      %#v\n", ages)
		fmt.Printf("ages's type   :   %T\n", ages)

		fmt.Printf("len(ages) :  %d\n", len(ages))
		fmt.Println("ages[0]:   ", ages[0])
		fmt.Println("ages[1]:   ", ages[1])
		fmt.Println("ages[len(ages)-1] : ", ages[len(ages)-1])

		ages[0] = 6
		ages[1] = -3
		fmt.Println("ages[0]      :    ", ages[0])
		fmt.Println("ages[1]      :    ", ages[1])

	}
}
