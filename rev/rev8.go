// 3. 특정  값을 설정한 컨텍스트 (rev8.go)

package main

import (
	"context"
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup // 하나의 고루틴을 대기하기 위해 사용

func main() {
	fmt.Println("컨텍스트에 대해 알아보자")

	wg2.Add(1) // 하나의 고루틴이 완료되기를 기다리도록 설정

	ctx := context.WithValue(context.Background(), "num", 9) // "num"이라는 키와 값 9를 저장한 새로운 컨텍스트 ctx를 생성

	go Square(ctx) // Square라는 함수를 새로운 고루틴에서 실행

	wg2.Wait() // Square 함수가 실행을 완료하고 wg.Done()을 호출할 때까지 대기 (WaitGroup의 카운터가 0이 될 때까지)
}

func Square(ctx context.Context) {
	if v := ctx.Value("num"); v != nil { // 컨텍스트에서 "num"라는 키로 저장된 값을 가져온다 -> interface{} 타입 반환
		// v가 nil(유효하지 않은 값)이면, "num" 키로 저장된 값이 없다는 뜻 -> 조건문 내부의 코드가 실행 X
		n := v.(int) // v는 interface{} 타입, v를 int 타입으로 변환
		fmt.Printf("Square : %d", n*n)
	}
	wg2.Done() // 작업이 완료되었음을 알리기 위해 -> WaitGroup의 카운터를 1 줄인다
}

/* if 조건문 사이에 ; 은 왜 쓰지?
- Go 언어에서 if 조건문 사이에 세미콜론(;)을 사용하는 이유는: if 문에서 변수를 선언하거나 초기화할 수 있도록 하기 위해서
- if 문 내에서 변수를 선언하고 동시에 그 변수를 조건에 사용할 수 있게 해준다.

1. 초기화 부분: v := ctx.Value("num")
- if 문에서 변수 v를 선언하고, ctx.Value("num")의 결과를 할당

2. 조건 부분: v != nil
- v가 nil이 아닌지 확인하는 조건

=> v := ctx.Value("num")를 먼저 실행하여 변수 v에 값을 할당하고,
그런 다음 v != nil 조건을 평가하여,
이 조건이 참(true)인 경우에만 if 블록 내부의 코드를 실행
*/

/* 키-값으로 저장하는데 맵 구조인가?
Go의 context.WithValue 함수는 특정한 키-값 쌍을 컨텍스트에 저장할 수 있도록 해준다
하지만 이 키-값 쌍은 맵(map)과 같은 자료 구조로 저장되는 것은 아니다!
대신, 컨텍스트는 키-값을 계층적으로 저장하는 방식으로 설계되어 있다.
*/
