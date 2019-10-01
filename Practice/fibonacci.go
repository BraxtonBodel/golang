package main

import "fmt"

func fibonacciFunc() {
	var n int
	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Print("Ingrese numero de terminos : ")
	fmt.Scan(&n)
	fmt.Println("Serie de fibonacci: ")
	for i := 1; i < n; i++ {
		if i == 1 {
			fmt.Println(" ", t1)
			continue
		}

		if i == 2 {
			fmt.Println(" ", t2)
			continue
		}

		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Println(" ", nextTerm)

	}

}
