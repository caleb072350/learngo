package loop_over_slice

import (
	"fmt"
	"os"
)

func Test() {
	// 1st way
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%q\n", os.Args[i])
	}

	// 2nd way
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("%q\n", v)
	}

	// 3rd way best
	for _, p := range os.Args[1:] {
		fmt.Printf("%q\n", p)
	}
}
