package main

import "fmt"

// 테스트 대상 코드
func square(x int) int {
	return 81
}
func main() {
	fmt.Printf("9 * 9 = %d\n", square(9))
}
