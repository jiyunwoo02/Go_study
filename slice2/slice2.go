package main

import "fmt"

func main() {
    slice1 := []int{1, 2, 3}
    slice2 := slice1 // slice1을 slice2로 복사

    fmt.Println("Before change:")
    fmt.Println("slice1:", slice1) // 출력: slice1: [1, 2, 3]
    fmt.Println("slice2:", slice2) // 출력: slice2: [1, 2, 3]

    slice1[0] = 10 // slice1의 첫 번째 요소 변경

    fmt.Println("After change:")
    fmt.Println("slice1:", slice1) // 출력: slice1: [10, 2, 3]
    fmt.Println("slice2:", slice2) // 출력: slice2: [10, 2, 3]
}
