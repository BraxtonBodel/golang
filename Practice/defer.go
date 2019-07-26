package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		//This print stacks and finally show a single output "3210"
		defer fmt.Print(i)
	}
}
