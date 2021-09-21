package loop_rand

import (
	"fmt"
	"math/rand"
	"time"
)

func TestRand() {
	rand.Seed(time.Now().UnixNano())

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
