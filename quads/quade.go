package quads

import (
	"fmt"
)

func QuadE(x, y int) {
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
			fmt.Print("C")
			for i := 0; i < x-2; i++ {
				fmt.Print("B")
			}
			if x > 1 {
				fmt.Print("A")
			}
			fmt.Println()
		}
	}
}

func quade() {
	QuadE(5, 3)
	QuadE(5, 1)
	QuadE(1, 1)
	QuadE(1, 5)
}
