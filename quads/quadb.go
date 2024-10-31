package quads

import (
	"fmt"
)

func QuadB(x, y int) {
	if x > 0 && y > 0 {

		fmt.Print("/")
		for i := 0; i < x-2; i++ {
			fmt.Print("*")
		}
		if x > 1 {
			fmt.Print("\\")
		}
		fmt.Println()

		for i := 0; i < y-2; i++ {
			fmt.Print("*")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			if x > 1 {
				fmt.Print("*")
			}
			fmt.Println()
		}

		if y > 1 {
			fmt.Print("\\")
			for i := 0; i < x-2; i++ {
				fmt.Print("*")
			}
			if x > 1 {
				fmt.Print("/")
			}
			fmt.Println()
		}
	}
}

func quadb() {
	QuadB(5, 3)
	QuadB(5, 1)
	QuadB(1, 1)
	QuadB(1, 5)
}
