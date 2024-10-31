package quads

import (
	"fmt"
)

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			fmt.Println("o")
			return
		}
		if x == 1 {
			fmt.Print("o\n")
			for i := 0; i < y-2; i++ {
				fmt.Println("|")
			}
			fmt.Println("o")
			return
		}
		fmt.Print("o")
		for i := 0; i < x-2; i++ {
			fmt.Print("-")
		}
		fmt.Println("o")

		for i := 0; i < y-2; i++ {
			fmt.Print("|")
			for j := 0; j < x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Println("|")
		}
		if y > 1 {
			fmt.Print("o")
			for i := 0; i < x-2; i++ {
				fmt.Print("-")
			}
			fmt.Println("o")
		}
	}
}

func quada() {
	QuadA(5, 3)
	QuadA(5, 1)
	QuadA(1, 1)
	QuadA(1, 5)
}
