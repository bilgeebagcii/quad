package quads

import (
	"fmt"
)

func QuadD(x, y int) {
	if x > 0 && y > 0 {

		fmt.Print("A")
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Print("C")
		}
		fmt.Println()

		for i := 0; i < y-2; i++ {
			fmt.Print("B")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			if x > 1 {
				fmt.Print("B")
			}
			fmt.Println()
		}

		if y > 1 {
			fmt.Print("A")
			for i := 0; i < x-2; i++ {
				fmt.Print("B")
			}
			if x > 1 {
				fmt.Print("C")
			}
			fmt.Println()
		}
	}
}

func quadd() {
	QuadD(5, 3)
	QuadD(5, 1)
	QuadD(1, 1)
	QuadD(1, 5)
}
