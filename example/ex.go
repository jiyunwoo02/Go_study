/*
2) 2를 Add하고 3개의 고루틴을 돌리면? 대부분 Panic
: 경우에 따라 에러가 발생하지 않고 잘 실행되는 경우도 가끔 있지만,
대부분 negative WaitGroup counter panic이 발생한다.
WaitGroup의 Add 값이 2인데, 3개의 고루틴이 실행되어 Done()을 3번 호출하면,
WaitGroup의 카운터는 0에서 -1로 내려가면서 panic이 발생한다.
*/

/*
1. Add(delta int): WaitGroup의 카운터에 값 추가, 일반적으로 고루틴을 시작하기 전에 호출되며, 시작할 고루틴의 수를 나타낸다.
2. Done(): 고루틴이 종료될 때마다 WaitGroup의 카운터를 감소시킨다.
3. Wait(): WaitGroup의 카운터가 0이 될 때까지 기다린다. 주로 모든 고루틴이 종료될 때까지 메인 고루틴이 기다리는 데 사용된다.

Go의 고루틴은 언제 시작되고 종료될지 정확히 예측하기 어렵다.
WaitGroup은 내부적으로 카운터를 관리하여 고루틴의 종료를 추적한다.
고루틴이 WaitGroup 카운터를 조작하는 시점과, Wait()가 호출되는 시점 사이의
-- 미묘한 타이밍 차이로 인해 때로는 에러가 발생하지 않을 수 있다!

*/

// 고루틴 중 일부가 Done을 호출하지 않도록 하면?
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2) // WaitGroup 카운터를 2로 설정
	var num int = 1

	// 고루틴 1
	go func() {
		defer wg.Done() // Done() 호출
		mu.Lock()
		fmt.Println("고루틴 1 작업 완료")
		num++ // num에 1을 더함
		mu.Unlock()
	}()

	// 고루틴 2
	go func() {
		defer wg.Done() // Done() 호출
		mu.Lock()
		fmt.Println("고루틴 2 작업 완료")
		num++ // num에 1을 더함
		mu.Unlock()
	}()

	// 고루틴 3
	go func() {
		// wg.Done() 호출 X
		mu.Lock()
		fmt.Println("고루틴 3 작업 완료")
		num++ // num에 1을 더함
		mu.Unlock()
	}()

	wg.Wait() // Wait()는 Done()이 2번 호출될 때까지 기다린다

	fmt.Println("2개의 고루틴 작업이 완료되었습니다.")
	fmt.Printf("num : %d", num)
}

/*
고루틴이 끝나기 전에 메인 함수가 종료되면 고루틴에서 수행하는 작업이 중단될 수 있다
-> WaitGroup을 사용하여 이러한 동기화 문제를 해결
-> WaitGroup을 사용하면 특정 고루틴이 완료된 후에 다른 작업을 실행
*/
