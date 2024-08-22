package main

import "fmt"

// 테스트 대상 코드
func Square(x int) int {
	return x * x
}
func main() {
	fmt.Printf("9 * 9 = %d\n", Square(9))
}
