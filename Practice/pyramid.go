package main

import "fmt"

func drawPyramid(maxRow int) {
	for i := 1; i <= maxRow; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
