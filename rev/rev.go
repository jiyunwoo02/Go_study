package main

import (
	"fmt"
)

/*
 3. 별칭 리시버 타입

type myInt int // 사용자 정의 별칭 타입

func (a myInt) add(b int) int { // myInt 별칭 타입을 리시버로 갖는 메서드
		return int(a) + b
	}
*/

/* 4. 인터페이스 선언
type Stringer interface {
	String() string // string 타입을 반환하는 String() 메서드
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕 나는 %d살 %s 라고 해", s.Age, s.Name) // Sprintf(출력할 형식, 출력값 또는 변수)
}
*/

/* 5. 빈 인터페이스 타입
func printVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("not supprted type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}
*/

func main() {

	fmt.Println("Go 1주차 스터디 리뷰")

	/* 5. 빈 인터페이스 타입
	printVal(10)          // v is int 10
	printVal(3.14)        // v is float64 3.140000
	printVal("Hello")     // v is string Hello
	printVal(Student{15}) // not supprted type: main.Student:{15}
	*/

	/* 4. 인터페이스 선언
	student := Student{"철수", 12}
	var stringer Stringer
	stringer = student
	fmt.Printf("%s", stringer.String())
	*/

	/* 3. 별칭 리시버 타입
	var a myInt = 10
	fmt.Println(a.add(30)) // 40
	var b int = 20
	fmt.Println(myInt(b).add(50)) // 70
	*/

	/* 1. 슬라이스 요소 추가
	var slice = []int{1, 2, 3} // [1 2 3]
	slice2 := append(slice, 4) // [1 2 3 4]
	fmt.Println(slice)
	fmt.Println(slice2)
	*/

	/* 2. 슬라이스에 여러 요소 추가
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1)) // slice1 :  [0 0 0] 3 5
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2)) // slice2 :  [0 0 0 4 5] 5 5

	slice1[1] = 100
	fmt.Println("slice1 두 번째 요소 값 변경 후")
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1)) // slice1 :  [0 100 0] 3 5
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2)) // slice2 :  [0 100 0 4 5] 5 5

	slice1 = append(slice1, 500)
	fmt.Println("slice1에 500 요소 추가")
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1)) // slice1 :  [0 100 0 500] 4 5
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2)) // slice2 :  [0 100 0 500 5] 5 5
	*/

}
