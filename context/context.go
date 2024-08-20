// Background() 는 최상단에서 사용이 되는 Context

package main

import (
	"context"
	"fmt"
	"time"
)

// 2. 작업 시간을 설정한 컨텍스트 (rev8.go)

func main() {
	// 생성된 시점으로부터 2초 후에 자동으로 취소되는 컨텍스트 생성
	//-> 즉, 2초가 지나면 ctx.Done() 채널이 닫히고, 이로 인해 타임아웃이 발생했음을 알릴 수 있다.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel() // 메인 함수가 종료될 때 컨텍스트 명시적으로 취소

	// 독립적인 고루틴에서 실행
	go dowork(ctx) // 작업을 수행하며, ctx 컨텍스트의 상태(컨텍스트가 취소되거나 시간이 초과되는) 모니터링

	// 메인 함수는 3초 동안 대기
	time.Sleep(3 * time.Second)
	// -> 2초 후에 컨텍스트가 취소되고, 3초 후에 메인 함수가 종료
}

func dowork(ctx context.Context) {
	select {
	case <-time.After(4 * time.Second): // <- 컨텍스트 시간보다 적게(예 : 1) 설정하면 이거 출력되는 거 확인
		// 4초 후에 작업 완료되면 메시지 출력
		fmt.Println("작업 완료")

	case <-ctx.Done():
		// 컨텍스트가 취소되거나 시간이 초과된 경우 메시지+취소 이유 출력
		fmt.Println("작업 취소: ", ctx.Err())
	}
}

// 2초가 지나면 컨텍스트가 타임아웃되고, 고루틴 내에서 작업이 취소된다
// why? 작업은 4초 후에 완료되도록 설정되어 있기 때문에

/*
2초 후: 컨텍스트 ctx가 취소되고, dowork 함수에서 "작업 취소: context deadline exceeded" 메시지 출력
3초 후: 메인 함수의 time.Sleep(3 * time.Second)가 완료되면서 메인 함수가 종료
-> 메인 함수가 종료되면 프로그램이 종료
*/

// --> 컨텍스트가 타임아웃되기 전에 작업이 끝나지 않아서 context deadline exceeded 오류 발생
// 만약 메인 함수가 time.Sleep을 호출하지 않고 즉시 종료된다면, 프로그램은 메인 고루틴이 종료되면서 다른 고루틴도 강제 종료
// => dowork 고루틴은 충분한 시간을 가지지 못하고 종료되기 때문에, dowork 함수의 출력 메시지가 나타나지 X
