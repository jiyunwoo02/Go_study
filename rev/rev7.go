/* < Context 세부적인 사용 예 >

context 패키지에서 제공하는 기능
컨텍스트는 작업을 지시할 때 작업 취소나 작업 시간 등을 설정할 수 있는 작업 명세서 역할을 한다.

새로운 고루틴으로 작업을 시작할 때 일정 시간 동안만 작업을 지시하거나, 외부에서 작업을 취소할 때 사용
또한 작업 설정에 관한 데이터 전달 가능

1. 작업 취소가 가능한 컨텍스트 (rev7.go)
: 작업자에게 컨텍스트를 만들어서 전달하면, 작업을 지시한 지시자가 원할 때 작업 취소를 알릴 수 있다

2. 작업 시간을 설정한 컨텍스트 (context.go)
: 일정한 시간 동안만 작업을 지시할 수 있다

3. 특정 값을 설정한 컨텍스트 (rev8.go)
: 때론 작업자에게 작업을 지시할 때 별도 지시사항을 추가하고 싶을 수 있다
: 컨텍스트에 특정 키로 값을 읽어올 수 있도록 설정할 수 있다

*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // WaitGroup 선언: 여러 고루틴의 작업이 완료될 때까지 대기
// 한 번의 고루틴이 생성될 때마다 카운터 1씩 증가시킨다 (고루틴의 시작과 종료에만 관여)

func main() {
	fmt.Println("컨텍스트에 대해 알아보자")

	wg.Add(1) // WaitGroup 카운터를 1 증가 -> 메인 고루틴이 PrintSecond 고루틴의 작업이 완료될 때까지 대기하도록 설정

	ctx, cancel := context.WithCancel(context.Background()) // 컨텍스트 ctx와 취소 함수 cancel 생성

	// ctx는 PrintSecond 고루틴에서 작업 상태를 모니터링하는 데 사용
	// cancel은 메인 고루틴에서 작업을 취소할 때 호출

	go PrintSecond(ctx) // 매 1초마다 "Tick"을 출력하며, 컨텍스트가 취소될 때까지 계속 실행

	time.Sleep(5 * time.Second) // 메인 고루틴이 5초 동안 대기 -> 이 동안 PrintSecond 고루틴이 Tick 출력

	cancel() // 5초 후에 컨텍스트 ctx 취소 -> PrintSecond 고루틴에서 이를 감지하여 종료
	cancel()
	// -> PrintSecond 고루틴이 종료되면
	// -> wg.Done()이 호출되어 카운터가 감소하고,
	// -> 메인 고루틴은 wg.Wait()를 통과해 프로그램이 종료

	wg.Wait() // WaitGroup의 카운터가 0이 될 때까지 대기
	//time.Sleep(time.Second)
	//fmt.Println("메인 끝")
}

// PrintSecond: 매 초마다 "Tick"을 출력하는 고루틴
func PrintSecond(ctx context.Context) {
	tick := time.Tick(time.Second) // 1초마다 신호를 보내는 채널 생성 -> 1초마다 출력 (카운터는 변하지 X)
	for {
		select {
		case <-ctx.Done(): // 컨텍스트의 취소 신호를 받으면 고루틴 종료
			fmt.Println("컨텍스트를 종료합니다.")
			wg.Done() // WaitGroup의 카운터를 감소(-1)시켜 메인 고루틴이 Wait() 대기를 마치게 함
			return    // -> 프로그램 종료
		case <-tick: // 매 1초마다 "Tick"을 출력
			fmt.Println("Tick")
		}
	}
}

/*컨텍스트에 대해 알아보자
Tick
Tick
Tick
Tick
Tick
컨텍스트를 종료합니다.
*/

// 결과: 총 5번의 Tick이 출력되고 프로그램이 종료된다.
// -> 단지, 고루틴이 시작되었기 때문에 wg.Add(1)로 카운터를 증가시킨 상태에서 고루틴이 끝나길 기다리는 것
