package main

import "fmt"

func evenOrOdd() {
	fmt.Print("Digite un numero : ")
	var n int
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "es un numero par")
	} else {
		fmt.Println(n, "es un numero impar")
	}
}
