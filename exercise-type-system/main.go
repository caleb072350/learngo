package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	// %b verb prints bits
	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)

	// %02b prints 2 bits with leading zeros if any
	fmt.Printf("%02b = %d\n", 0, 0)
	fmt.Printf("%02b = %d\n", 1, 1)
	fmt.Printf("%02b = %d\n", 2, 2)
	fmt.Printf("%02b = %d\n", 3, 3)

	// %08b prints 8 bits with leading zeros if any
	fmt.Printf("%08b = %d\n", 1, 1)
	fmt.Printf("%08b = %d\n", 2, 2)
	fmt.Printf("%08b = %d\n", 4, 4)
	fmt.Printf("%08b = %d\n", 8, 8)
	fmt.Printf("%08b = %d\n", 16, 16)
	fmt.Printf("%08b = %d\n", 32, 32)
	fmt.Printf("%08b = %d\n", 64, 64)
	fmt.Printf("%08b = %d\n", 128, 128)

	i, _ := strconv.ParseInt("00010110", 2, 64)
	fmt.Println(i)

	var b byte
	b = 0
	fmt.Printf("%08b = %d\n", b, b)
	b = 255
	fmt.Printf("%08b = %d\n", b, b)

	// you can find the limits of numerics types in the math package
	fmt.Println("int8     :", math.MinInt8, math.MaxInt8)
	fmt.Printf("%08b = %d\n", math.MinInt8, math.MinInt8)

	// memory costs
	fmt.Println("int8 : ", unsafe.Sizeof(int8(1)))
	fmt.Println("int64 : ", unsafe.Sizeof(int64(1)))

	// overflow
	var width uint8 = 255
	width++
	fmt.Println("width: ", width)

	h, _ := time.ParseDuration("4h30m")
	fmt.Println(h.Hours(), "hours")

	var m int64 = 2
	h *= time.Duration(m)
	fmt.Println(h)

	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))
}
