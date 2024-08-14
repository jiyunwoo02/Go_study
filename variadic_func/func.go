package main

import "fmt"

func sum(nums ...int) int{
	total := 0

	for _, num := range nums{
		total += num
	}
	return total
}

func main(){
	result1 := sum(1,2,3)
	fmt.Println("Sum : ", result1)

	result2 := sum(1,2,3,4)
	fmt.Println("Sum : ", result2)

	result3 := sum()
	fmt.Println("Sum : ", result3)
}