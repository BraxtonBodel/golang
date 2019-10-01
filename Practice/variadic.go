package main

import "fmt"

func calculation(word string, number ...int) (string, int) {
	//area := 1
	for _, value := range number {
		fmt.Println(value, number)
	}
	return "resultado", 10
}
