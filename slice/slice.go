package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println("Slice", s)
	fmt.Println("Length : ", len(s))
	fmt.Println("Capacity : ", cap(s))
	fmt.Println("Pointer : ", &s[0])
}
