/* < 4. time.Tick(), After() >

time.Tick()과 time.After()는 Go에서 시간 관련 작업을 수행할 때 유용하게 사용할 수 있는 기능
- 이들 각각은 주기적인 작업과 특정 시간 후의 작업을 처리할 때 사용

time.Tick()는 일정 시간 간격 주기로 신호를 보내주는 채널을 생성해서 반환하는 함수
- 이 함수가 반환한 채널에서 데이터를 읽어오면, 일정 시간 간격으로 현재 시각을 나타내는 Time 객체 반환

time.After()는 현재 시간 이후로 일정 시간 경과 후에 신호를 보내주는 채널을 생성해서 반환하는 함수
- 이 함수가 반환한 채널에서 데이터를 읽으면, 일정 시간 경과 후에 현재 시각을 나타내는 Time 객체 반환

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time 패키지의 Tick()와 After()에 대해 알아보자\n")

	// 1. 2초 간격의 틱 채널 생성
	tickInterval := 2 * time.Second     // 틱 간격을 2초로 설정
	tickChan := time.Tick(tickInterval) // 2초마다 틱을 발생시키는 '채널' 생성(unbuffered) -> 지정한 간격(2초)마다 값(현재 시각)을 송신

	// 2. 5초 후에 닫힐 채널 생성
	timeout := 5 * time.Second         // 타임아웃 시간 5초로 설정
	timeoutChan := time.After(timeout) // 5초 후에 단일 값을 보내는 '채널' 생성(unbuffered) -> 지정한 시간(5초) 후에 한 번만 값(현재 시각)을 전송하고 종료

	// select 문은 for 루프가 실행되는 동안 반복적으로 평가
	for { // select 문이 하나의 case를 처리하면, for 루프는 다시 반복을 시작하여 select 문을 다시 평가
		// 송신자는 time.Tick()과 time.After() 함수, 수신자: main 함수 내의 select 문
		select {
		case tick := <-tickChan: // tick : Tick()으로부터 생성된 채널(tickChan)에서 수신한 time.Time 값(현재 시각)을 저장하는 변수
			// 틱 이벤트를 2초마다 처리, 2초마다 아래 코드가 실행된다
			fmt.Println("틱 수신 시각 : ", tick) // => 코드 실행 후 2초 간격으로 두 번의 틱 이벤트가 발생

		case timeoutTime := <-timeoutChan: // After()으로부터 생성된 채널(timeoutChan)에서 값(현재 시각)을 수신하고 값을 저장하는 데 사용되는 변수
			// 5초 후 타임아웃 이벤트 처리
			fmt.Println("\n타임아웃 시각:", timeoutTime) // => 프로그램 실행 후 5초 후에 타임아웃 이벤트가 발생
			fmt.Println("종료합니다.")
			return // for 루프를 종료하고 main() 함수의 실행 종료

		default: // 모든 case가 준비되지 않았을 때 실행
			// default 문이 없으면 -> select 문은 계속해서 반복 평가되며, CPU 자원을 낭비
			// CPU 과다 사용을 방지하기 위해 짧은 시간 대기
			time.Sleep(100 * time.Millisecond) // 100밀리초 동안 대기

			//  Tick()은 2초와 4초에 실행되고, After()는 5초에 실행
		}
	}
}

/*time 패키지의 Tick()와 After()에 대해 알아보자

틱 수신 시각 :  2024-08-20 14:32:11.5205943 +0900 KST m=+2.006829801
틱 수신 시각 :  2024-08-20 14:32:13.5219744 +0900 KST m=+4.008209901

타임아웃 시각: 2024-08-20 14:32:14.5282725 +0900 KST m=+5.014508001
종료합니다.
*/

// 3초로 틱 해두면 3초 한 번 출력되고 종료된다

/*
for 루프는 무한 루프이므로, select 문은 반복적으로 실행된다
- tickChan이 준비되어 있는 동안, select 문은 tickChan의 값을 수신하고, tickChan의 case를 실행
- 이 과정은 timeoutChan에서 값이 수신되기 전까지 계속 반복된다.

프로그램 실행 후 약 2초 후에 첫 번째 틱 이벤트가 발생하고, 약 4초 후에 두 번째 틱 이벤트가 발생.
약 5초 후에 timeoutChan에서 값이 수신되면서 타임아웃 이벤트가 발생하고, 프로그램이 종료

1. 0~2초: 프로그램이 시작된 후, 2초마다 tickChan에서 값이 발생, tick 값이 출력
2. 2초: 첫 번째 tickChan 값이 출력
3. 4초: 두 번째 tickChan 값이 출력
4. 5초: timeoutChan에서 값이 발생하고, timeoutTime 값이 출력되며, 프로그램이 종료

Nanosecond (ns): 1초의 10억 분의 1. 1 * time.Nanosecond
Microsecond (µs): 1초의 백만 분의 1. 1 * time.Microsecond
Millisecond (ms): 1초의 천 분의 1. 1 * time.Millisecond
Second (s): 기본 시간 단위. 1 * time.Second
Minute (m): 60초. 1 * time.Minute
Hour (h): 3600초. 1 * time.Hour
Day (d): 86400초 (24시간). Go의 표준 라이브러리에서는 직접적으로 사용되는 상수는 없지만, 직접 계산할 수 있다.

*/
