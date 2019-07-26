package main

import "fmt"

func openFile(src, path string) (sucess, err int64) {
	sucess, err = 12112, 23121
	return sucess, err
}

func main() {
	var path, thin string
	var isFine, panic = openFile(path, thin)
	fmt.Print("Type : %T Value: %v", isFine, isFine)
	fmt.Print("Type : %T Value: %v", panic, panic)
}
