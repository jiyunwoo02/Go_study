/*
1. 메인 고루틴이 Wait()를 호출하기 전에 모든 고루틴이 종료된 경우:
만약 3개의 고루틴 중 1개의 고루틴이 Done()을 호출하고 종료된 시점에
- 메인 고루틴이 Wait()를 호출한다면,
그 시점에서 WaitGroup의 카운터가 이미 0이 되어 있을 수 있다.
- 이 경우 Wait()는 즉시 반환되므로 에러가 발생하지 않는다.

2. 비정상적으로 빠른 고루틴 실행:
만약 고루틴이 매우 빠르게 실행되고 종료된다면,
- 고루틴이 Done()을 호출하기 전에 Wait()가 호출되지 않아서
- negative counter가 발생하기 전에 프로그램이 정상적으로 종료될 수 있다.

3. 경쟁 조건 (Race Condition):
경쟁 조건이 발생하면 고루틴들이 Done()을 호출하는 타이밍이 겹치면서
- 비정상적인 상황이 발생할 수 있다.
- 이로 인해 때로는 패닉이 발생하지 않을 수 있다.
- 이는 프로그램의 실행 순서가 달라질 수 있음을 의미
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add 값을 2로 설정합니다.
	wg.Add(2)

	//var num int = 1

	// 고루틴 1
	go func() {
		defer wg.Done()             // Done() 호출
		time.Sleep(1 * time.Second) // 잠시 대기
		fmt.Println("고루틴 1 작업 완료")
		//num++
	}()

	// 고루틴 2
	go func() {
		defer wg.Done()             // Done() 호출
		time.Sleep(1 * time.Second) // 잠시 대기
		fmt.Println("고루틴 2 작업 완료")
		//num++
	}()

	// 고루틴 3
	go func() {
		time.Sleep(2 * time.Second) // 잠시 대기
		fmt.Println("고루틴 3 작업 완료")
		//num++
		// Done() 호출 X
	}()

	// 메인 고루틴 대기
	//time.Sleep(5 * time.Second)

	// 이 시점에서 Done()이 2번 호출되었고, 1번 더 호출될 수 있음
	wg.Done() // 의도적으로 추가 Done 호출 (3번째 Done)

	// Wait()는 Done()이 2번 호출될 때까지 기다림
	wg.Wait()

	fmt.Println("작업이 완료되었습니다.")
	//fmt.Printf("num : %d\n", num)
}

// 패닉은 qait에서만
