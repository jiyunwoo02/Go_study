/* < 2. 채널 unbuffered, buffered 동작 차이 >

Go에서 채널(Channel)은 고루틴(goroutine) 간의 통신을 위해 사용되며, 데이터를 안전하게 주고받을 수 있게 해준다.

채널에는 두 가지 주요 유형이 있다:
- 버퍼링된 채널(Buffered Channel)
- 버퍼링되지 않은 채널(Unbuffered Channel)

1. 버퍼링되지 않은 채널: 송신자와 수신자가 동시에 준비되어야만 데이터를 주고받을 수 있는 채널
	- 동기식 통신: 송신자와 수신자가 동시에 동작해야 함!
	- 송신자(sender)가 데이터를 채널에 보내려고 하면, 그 데이터를 받는 수신자(receiver)가 있을 때까지 송신자는 블록된다(대기 상태).
	- 반대로, 수신자는 채널에서 데이터를 받으려고 할 때, 송신자가 데이터를 보낼 때까지 블록된다.

2. 버퍼링된 채널: 일정 크기의 버퍼(큐)를 가지며, 송신자가 수신자가 없어도 버퍼가 가득 차지 않았을 때 데이터를 채널에 보낼 수 있는 채널
	- 비동기식 통신: 송신자와 수신자가 동시에 동작할 필요 없음!
	- 버퍼링된 채널의 경우, 송신자는 채널이 가득 찰 때까지 데이터를 보내고 바로 반환된다.
	- 수신자는 채널이 비어 있을 때까지 데이터를 받을 수 있다.

=> 동기화가 필요한 경우에는 unbuffered channel을, 비동기적으로 데이터를 주고받고자 할 때는 buffered channel을 사용

일반적으로 채널을 생성하면 크기가 0인 채널이 만들어진다 (unbuffered)
- 크기가 0이라는 뜻은 채널에 들어온 데이터를 담아둘 곳이 없다는 얘기가 된다

ex) 예를 들어 택배 기사가 택배를 전달하는데
- 택배를 담아둘 곳이 없으면 수신자가 와서 가져갈 때까지 송신자가 택배를 들고 기다려야 한다
- 하지만 택배 보관함이 있고 보관함에 공간이 남아 있으면 송신자가 보관함에 택배를 넣고 다른 일을 할 수 있다

채널 크기가 0이다 = '택배 보관함이 없는 경우'
- 즉, 데이터를 넣을 때 보관할 곳이 없기 때문에 데이터를 빼갈 때까지 대기한다
- 어느 수신자도 데이터를 빼가지 않으면, 송신자는 영원히 대기하게 되어 deadlock 메시지를 출력하고 프로그램이 강제 종료된다

*버퍼: 내부에 데이터를 보관할 수 있는 메모리 영역
보관함을 가지고 있는 채널 = 버퍼를 가진 채널 -> make() 함수에서 버퍼 크기를 적어준다 (buffered)

버퍼가 다 차면, 버퍼가 없을 때와 마찬가지로 보관함에 빈 자리가 생길 때까지 대기한다
그래서 데이터를 제때 빼가지 않으면, 버퍼가 없을 때처럼 대기

*/

package main

import (
	"fmt"
	"time"
)

func main() { // 메인 고루틴[수신자] 실행: 각 고루틴은 채널을 통해 메인 고루틴으로 데이터를 보낸다

	fmt.Println("Channel에 대해 알아보자\n")

	// Unbuffered Channel
	unbufferedChan := make(chan string) // 버퍼 크기가 0인 채널, 채널에 데이터 보내려면 수신자가 즉시 있어야 한다
	// -> 데이터를 보내는 고루틴은 데이터가 채널에 수신될 때까지 블록된다

	// Buffered Channel with a buffer size of 2
	bufferedChan := make(chan string, 2) // 버퍼 크기가 2인 채널, 최대 2개의 데이터를 수신자가 받지 않아도 저장 가능
	// 2개 데이터 저장된 이후에 송신자는 더 이상 데이터 보낼 수 없으며, 수신자가 데이터를 받을 때까지 블록된다.
	// -> 두 개의 값을 저장할 수 있으며, 송신자는 버퍼가 가득 차기 전까지 블록되지 않는다.

	// Function to Unbuffered Channel[송신자]
	go func() {
		fmt.Println("Goroutine 1: unbuffered 채널에 데이터 보내기") // 데이터 보내기 시작
		unbufferedChan <- "unbuffered 채널 메시지"              // 메인 고루틴이 이 채널에서 데이터 받을 준비가 될 때까지 고루틴1은 멈춘다
		fmt.Println("Goroutine 1: unbuffered channel에 데이터 보내짐")
	}()

	// Function to Buffered Channel[송신자]
	go func() {
		fmt.Println("Goroutine 2: buffered 채널에 데이터 보내기")
		bufferedChan <- "buffered 채널 메시지1"
		bufferedChan <- "buffered 채널 메시지2" // 두 개의 메시지가 연속으로 전송되어 모두 버퍼에 저장되고, 고루틴2는 블록되지 않는다 (수신자가 없어도 송신 가능)
		fmt.Println("Goroutine 2: buffered 채널에 데이터 보내짐")
	}()

	time.Sleep(3 * time.Second) // 메인 고루틴을 3초 동안 지연
	fmt.Println("\n-- 메인 고루틴 3초간 대기 완료!\n")
	// time.Second는 1초, time.Sleep 함수는 인자로 받은 시간을 동안 현재 고루틴을 "수면 상태"로 만들어, 다른 작업을 할 수 없게 만든다.

	// Receiving from Unbuffered Channel
	// 3초 후 메인 고루틴이 깨어나 데이터 수신
	fmt.Println("Main: unbuffered 채널로부터 데이터 수신")
	fmt.Println("Main: 수신 완료 - ", <-unbufferedChan) // 메인 고루틴이 실제로 데이터를 수신

	// Receiving from Buffered Channel
	// 메인 고루틴이 bufferedChan으로부터 두 개의 데이터를 수신한다. (고루틴2는 이미 데이터 전송 완료한 상태)
	fmt.Println("Main: buffered 채널로부터 데이터 수신")
	fmt.Println("Main: 수신 완료 - ", <-bufferedChan)
	fmt.Println("Main: 수신 완료 - ", <-bufferedChan)

	/*Channel에 대해 알아보자

	Goroutine 1: unbuffered 채널에 데이터 보내기
	Goroutine 2: buffered 채널에 데이터 보내기
	Goroutine 2: buffered 채널에 데이터 보내짐

	-- 메인 고루틴 3초간 대기 완료!

	Main: unbuffered 채널로부터 데이터 수신
	Main: 수신 완료 -  unbuffered 채널 메시지
	Main: buffered 채널로부터 데이터 수신
	Goroutine 1: unbuffered channel에 데이터 보내짐
	Main: 수신 완료 -  buffered 채널 메시지1
	Main: 수신 완료 -  buffered 채널 메시지2
	*/

	// 각 고루틴은 데이터를 채널을 통해 메인 고루틴으로 송신
	// 메인 고루틴은 이 데이터를 수신하고 출력하는 역할
}

/*
1. 메인 고루틴이 시작되고, Unbuffered Channel과 Buffered Channel을 각각 생성
2. 고루틴 1이 unbufferedChan에 데이터를 보내려 하지만, 메인 고루틴이 수신할 준비가 되지 않았기 때문에 일시적으로 블록됨
3. 고루틴 2는 bufferedChan에 두 개의 데이터를 전송하고, 버퍼가 가득 차지 않았으므로 블록되지 않음
4. 메인 고루틴이 3초 동안 대기한 후 깨어나서, 각 채널에서 데이터를 수신
5. unbufferedChan에서 데이터를 수신하면 고루틴 1이 블록 상태에서 해제
6. bufferedChan에서 두 개의 데이터를 차례로 수신
*/

/*
이 프로그램에는 총 3개의 고루틴이 있다: (매번 순서 랜덤)

1. 메인 고루틴: 프로그램이 시작되면 main 함수가 실행되며, 이것이 메인 고루틴
2. 첫 번째 고루틴: unbufferedChan에 데이터를 보내는 고루틴
3. 두 번째 고루틴: bufferedChan에 데이터를 보내는 고루틴

>> unbuffered channel 생성: 버퍼 크기가 0인 채널이기 때문에, 송신자가 데이터를 보내려면 즉시 수신자가 존재해야 함. 송신자가 수신자를 기다리는 방식

>> buffered channel 생성: 버퍼 크기가 2인 채널이기 때문에, 수신자가 없더라도 최대 2개의 데이터를 버퍼에 저장할 수 있다.
- 송신자는 버퍼가 가득 차기 전까지 기다리지 않고 데이터를 보낼 수 있다.
- 즉, 메인 고루틴이 즉시 데이터를 수신하지 않아도, 송신자는 버퍼가 가득 차기 전까지 데이터를 계속 보낼 수 있다.
-- Buffered 채널은 메인 고루틴이 즉시 데이터를 수신할 필요는 없지만, 결국에는 데이터를 수신해야 한다. 그렇지 않으면 버퍼가 가득 차고, 송신자는 블록된다.

[1] 첫 번째 고루틴은 unbufferedChan에 데이터를 보내려고 한다.
그러나 Unbuffered 채널이기 때문에, 메인 고루틴이 데이터를 수신할 때까지 이 고루틴은 unbufferedChan <- "unbuffered 채널 메시지"에서 블록된다

[2] 두 번째 고루틴은 bufferedChan에 두 개의 데이터를 연속으로 보낸다.
bufferedChan은 버퍼 크기가 2이기 때문에, 이 두 개의 데이터는 버퍼에 저장되고, 이 고루틴은 블록되지 않고 계속해서 실행된다.

[3] 메인 고루틴은 프로그램이 시작되면 실행되는 기본 고루틴
- 메인 고루틴은 다른 고루틴을 실행하는 역할을 수행
- 다른 고루틴들이 데이터를 주고받을 수 있도록 채널을 통해 데이터를 수신하는 역할
-- 만약 메인 고루틴이 없었다면, unbufferedChan에서 데이터를 보내려는 고루틴이 블록된 채로 남아 프로그램이 멈추게 되었을 것
-- 메인 고루틴이 데이터를 수신함으로써 프로그램이 정상적으로 종료될 수 있게 된다.

[4] time.Sleep(3 * time.Second)는 메인 고루틴을 3초 동안 멈추게 한다.
- 이 동안 다른 고루틴들이 작업을 수행할 수 있게 시간을 주는 것

[5] 메인 고루틴은 unbufferedChan에서 데이터를 수신
- 이 작업이 수행되기 전까지는 unbufferedChan에 데이터를 보내려는 첫 번째 고루틴이 블록 상태였다.
이제 메인 고루틴이 데이터를 수신하므로, 첫 번째 고루틴은 블록 상태에서 벗어나게 된다.

[6] 메인 고루틴은 bufferedChan에서 두 개의 데이터를 순차적으로 수신
- 이 채널의 버퍼는 이미 데이터가 채워져 있으므로, 메인 고루틴은 블록되지 않고 즉시 데이터를 받을 수 있다.
*/
