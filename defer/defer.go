package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("example1.txt")
	// -> Error opening file:  open example1.txt: The system cannot find the file specified. 오류 발생!
	file, err := os.Open("example1.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	fmt.Println("file opened successfully")
}
