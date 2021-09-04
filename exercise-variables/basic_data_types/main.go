package main

import "fmt"

func main() {
	fmt.Printf("%d %d\n", 2, 3)
	fmt.Printf("%4.2f %4.2f\n", 3.14, 4.35)
	fmt.Printf("%t\n", true)

	name := "Caleb"
	fmt.Printf("%s\n", name)

	test := "中国"
	fmt.Printf("%s\n", test)

	fmt.Printf("%X\n", 15)

	// 变量声明
	var speed int
	var counter int
	var nCPU int
	fmt.Println(speed, counter, nCPU)

	var heat float64
	var ratio float64
	var degree float64
	fmt.Println(heat, ratio, degree)

	var off bool
	var valid bool
	var closed bool
	fmt.Println(off, valid, closed)

	var msg string
	var came string
	var text string
	fmt.Println(msg, came, text)

	var (
		spd  int
		sf   float64
		bof  bool
		asfl string
	)

	fmt.Println(spd)
	fmt.Println(sf)
	fmt.Println(bof)
	fmt.Println(asfl)
}
