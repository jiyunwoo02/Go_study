package main

import "fmt"

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice1 []int = array[1:5]           // 1. 배열 슬라이싱
	var slice2 []int = slice1[1:8:9]        // 2. 슬라이스 슬라이싱
	var slice3 []int = make([]int, 5)       // 3. make()
	var slice4 []int = make([]int, 0)       // 4. 길이 0인 슬라이스
	var slice5 []int = []int{1, 2, 3, 4, 5} // 5. 초기화
	var slice6 []int                        // 6. 기본값은 nil

	fmt.Println("slice1 : ", slice1) // slice1 :  [2 3 4 5]
	fmt.Println("slice2 : ", slice2) // slice2 :  [3 4 5 6 7 8 9], cap은 9-1=8
	fmt.Println("slice3 : ", slice3) // slice3 :  [0 0 0 0 0] -> 모든 요소는 해당 타입의 기본값으로 초기화된다! (int 타입 기본값=0)

	fmt.Println("slice4 : ", slice4) // slice4 :  [], 초기화된 상태로 간주!, 길이와 용량이 모두 0
	// 길이와 용량이 모두 0으로 설정되었기 때문에, 실제 데이터를 저장하기 위한 메모리 공간은 할당되지 않는다
	// 내부 포인터는 실제로 데이터를 가리키지 않을 수도 있다 (nil 포인터일 수 있음)
	if slice4 != nil {
		fmt.Println("slice4 is not nil")
	} // slice4 is not nil

	fmt.Println("slice5 : ", slice5) // slice5 :  [1 2 3 4 5]

	fmt.Println("slice6 : ", slice6) // slice6 :  [], 변수 초기화 X이기에 nil로 설정, 내부 포인터는 nil을 가리키고, 길이와 용량은 0
	if slice6 == nil {
		fmt.Println("slice6 is nil")
	} // slice6 is nil

}
