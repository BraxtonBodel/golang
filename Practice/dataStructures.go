package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(items []int) {
	var length = len(items)
	var sorted = false

	for !sorted {
		swapped := false
		for i := 0; i < length-1; i++ {
			fmt.Println(items[i])
			swapped = true
		}

		if !swapped {
			sorted = true
		}
	}
	fmt.Println(length)
}

//Here we create an array fulled with random numbers and then we return the array
func generateSlices(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
