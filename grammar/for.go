package main

import "fmt"

func main() {

	// 1. 문자열
	str := "hello 월드"
	for i, c := range str {
		fmt.Printf("1) 인덱스 : %d, 문자열: %c\n", i, c)
	}
	/* <결과>
	인덱스 : 0, 문자열: h
	인덱스 : 1, 문자열: e
	인덱스 : 2, 문자열: l
	인덱스 : 3, 문자열: l
	인덱스 : 4, 문자열: o
	인덱스 : 5, 문자열:
	인덱스 : 6, 문자열: 월
	인덱스 : 9, 문자열: 드
	*/

	// 2. 슬라이스
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Printf("인덱스 : %d, 요소 : %d\n", i, v)
	}
	/* <결과>
	인덱스 : 0, 요소 : 1
	인덱스 : 1, 요소 : 2
	인덱스 : 2, 요소 : 3
	인덱스 : 3, 요소 : 4
	인덱스 : 4, 요소 : 5
	*/

	// 3. 맵
	m := map[string]int{"aaa": 1, "bbb": 2, "ccc": 3}
	for k, v := range m {
		fmt.Printf("키 : %s, 값 : %d\n", k, v)
	}
	/* <결과>
	키 : bbb, 값 : 2
	키 : ccc, 값 : 3
	키 : aaa, 값 : 1
	*/

	// 4. 채널
	ch := make(chan int)
	go func() {
		for i := 1; i <= 4; i++ {
			ch <- i
		}
		close(ch) // 채널 닫기: for range 루프에 더 이상 받을 값이 없다는 것을 알림
	}()
	// 채널에서 값을 받아서 출력
	for num := range ch {
		fmt.Printf("값 : %d\n", num)
	}
	/* 결과>
	값 : 1
	값 : 2
	값 : 3
	값 : 4
	*/
}
