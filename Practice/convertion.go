package main

import (
	"fmt"
	"strconv"
)

func convertion() {
	stringVariable := "100"
	intVariable, _ := strconv.Atoi(stringVariable)
	fmt.Printf("%T\n", intVariable)
}
