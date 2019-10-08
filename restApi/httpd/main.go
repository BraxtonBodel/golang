package main

import (
	"findi/platform/findi"
	"fmt"
)

func main() {
	// r := gin.Default()

	// r.GET("/ping", handler.PingGet())

	// r.Run()

	find := findi.New()
	fmt.Println(find)
	find.Add(findi.Item{"Hello", "hello world"})
	fmt.Println(find)
}
