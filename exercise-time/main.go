package main

import (
	"fmt"
	"path"
	"time"
)

func main() {
	var (
		speed     int
		now       time.Time
		prevSpeed int
	)

	speed, now = 100, time.Now()

	fmt.Println(speed, now)

	fmt.Println(now.Format(time.Kitchen))

	prevSpeed = 50

	speed, prevSpeed = prevSpeed, speed

	fmt.Println(speed, prevSpeed)

	var dir, file string
	dir, file = path.Split("/test/css/main.css")

	fmt.Println("dir:", dir)
	fmt.Println("file:", file)

}
