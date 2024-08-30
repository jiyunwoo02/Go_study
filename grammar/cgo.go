package main

/*
#include <stdlib.h>
*/
import "C" // Cgo 기능 사용
import "fmt"

func Random() int {
	return int(C.rand()) // stdlib.h에 정의된 rand() 함수 호출
	// C의 rand() 함수에서 반환된 값을 Go의 int 타입으로 변환하여 반환

	// rand() 함수를 호출하여 0과 RAND_MAX 사이의 난수를 생성하고 반환
	// RAND_MAX는 시스템에 따라 다를 수 있지만, 보통 32767 (2^15 - 1)
}

func Seed(i int) { // 난수 생성기에 시드 값 제공
	C.srand(C.uint(i))
	// i는 Go의 int 타입으로 제공되며, C.uint(i)를 통해 C의 unsigned int 타입으로 변환

	// srand() 함수: 난수 생성기의 시작점을 설정하는 데 사용
	// 시드 값이 같으면 rand() 함수는 동일한 난수 시퀀스를 반복적으로 생성
}

func main() {
	Seed(1)
	fmt.Println(Random())
}
