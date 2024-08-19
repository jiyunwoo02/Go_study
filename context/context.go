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

	defer cancel() // 메인 함수가 종료될 때 컨텍스트 취소

	// 고루틴에서 작업 수행
	go dowork(ctx) // 작업을 수행하며, ctx 컨텍스트의 상태를 모니터링

	// 메인 함수는 3초 동안 대기
	time.Sleep(3 * time.Second) // 고루틴의 작업이 완료되거나, 컨텍스트가 취소될 때까지 대기
	// -> 이 대기 시간은 컨텍스트의 타임아웃보다 길게 설정되어, 컨텍스트가 먼저 타임아웃되도록 유도
}

func dowork(ctx context.Context) {
	select {
	case <-time.After(4 * time.Second): // 컨텍스트 시간보다 적게(예 : 1) 설정하면 이거 출력되는 거 확인
		// 5초 후에 작업 완료되면 메시지 출력
		fmt.Println("작업 완료")

	case <-ctx.Done():
		// 컨텍스트가 취소되거나 시간이 초과된 경우 메시지+취소 이유 출력
		fmt.Println("작업 취소: ", ctx.Err())
	}
}

// 2초가 지나면 컨텍스트가 타임아웃되고, 고루틴 내에서 작업이 취소된다
// why? 작업은 5초 후에 완료되도록 설정되어 있기 때문에
// --> 컨텍스트가 타임아웃되기 전에 작업이 끝나지 않아서 context deadline exceeded 오류 발생
