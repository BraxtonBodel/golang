package main

import "fmt"

func drawPyramid(maxRow int) {
	for i := 1; i <= maxRow; i++ {
		if i < 7 {
			for j := 1; j < i; j++ {
				fmt.Print("o ")
			}
		} else {
			x := (i - 6) * 2
			for j := i; j < x; j-- {
				fmt.Print(" o")
			}
		}
		fmt.Println("")
	}
}
