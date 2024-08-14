/* 2. 채널 unbuffered, buffered 동작 차이
Go에서 채널(Channel)은 고루틴(goroutine) 간의 통신을 위해 사용되는 데이터 파이프라인으로, 데이터를 안전하게 주고받을 수 있게 해준다.
채널에는 두 가지 주요 유형이 있다:
**버퍼링된 채널(Buffered Channel)**과 버퍼링되지 않은 채널(Unbuffered Channel)

1. 버퍼링되지 않은 채널은 송신자와 수신자가 동시에 준비되어야만 데이터를 주고받을 수 있는 채널
	- 동기식 통신: 송신자와 수신자가 동시에 동작해야 함
	- 송신자(sender)가 데이터를 채널에 보내려고 하면, 그 데이터를 받는 수신자(receiver)가 있을 때까지 송신자는 블록된다(대기 상태).
	- 반대로, 수신자는 채널에서 데이터를 받으려고 할 때, 송신자가 데이터를 보낼 때까지 블록된다.

2. 버퍼링된 채널은 일정 크기의 버퍼(큐)를 가지며, 송신자가 수신자가 없어도 버퍼가 가득 차지 않았을 때 데이터를 채널에 보낼 수 있는 채널
	- 비동기식 통신: 송신자와 수신자가 동시에 동작할 필요 없음
	- 버퍼링된 채널의 경우, 송신자는 채널이 가득 찰 때까지 데이터를 보내고 바로 반환된다.
	- 수신자는 채널이 비어 있을 때까지 데이터를 받을 수 있다.

=> Go에서 어떤 채널을 사용할지는 상황에 따라 다르며, 동기화가 필요한 경우에는 unbuffered channel을, 비동기적으로 데이터를 주고받고자 할 때는 buffered channel을 사용

일반적으로 채널을 생성하면 크기가 0인 채널이 만들어진다 (unbuffered)
크기가 0이라는 뜻은 채널에 들어온 데이터를 담아둘 곳이 없다는 얘기가 된다

예를 들어 택배 기사가 택배를 전달하는데
- 택배를 담아둘 곳이 없으면 수신자가 와서 가져갈 때까지 택배를 들고 기다려야 한다
- 하지만 택배 보관함이 있고 보관함에 공간이 남아 있으면 보관함에 택배를 넣고 다른 일을 할 수 있다

채널 크기가 0이다 = 택배 보관함이 없는 경우
- 즉 데이터를 넣을 때 보관할 곳이 없기 때문에 데이터를 빼갈 때까지 대기한다
- 어느 수신자도 데이터를 빼가지 않으면, 송신자는 영원히 대기하게 되어 deadlock 메시지를 출력하고 프로그램이 강제 종료된다

*버퍼: 내부에 데이터를 보관할 수 있는 메모리 영역
보관함을 가지고 있는 채널 = 버퍼를 가진 채널 -> make() 함수에서 뒤에 버퍼 크기를 적어준다 (buffered)

버퍼가 다 차면, 버퍼가 없을 때와 마찬가지로 보관함에 빈 자리가 생길 때까지 대기한다
그래서 데이터를 제때 빼가지 않으면, 버퍼가 없을 때처럼 대기

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered Channel
	unbufferedChan := make(chan string) // 데이터를 보내는 고루틴은 데이터가 채널에 수신될 때까지 블록된다

	// Buffered Channel with a buffer size of 2
	bufferedChan := make(chan string, 2) // 두 개의 값을 저장할 수 있으며, 송신자는 버퍼가 가득 차기 전까지 블록되지 않는다.

	// Function to demonstrate Unbuffered Channel
	go func() {
		fmt.Println("Goroutine 1: Sending data to unbuffered channel...")
		unbufferedChan <- "Message to unbuffered channel" // 데이터가 메인 고루틴에서 수신될 때까지 블록된다
		fmt.Println("Goroutine 1: Data sent to unbuffered channel.")
	}()

	// Function to demonstrate Buffered Channel
	go func() {
		fmt.Println("Goroutine 2: Sending data to buffered channel...")
		bufferedChan <- "Message 1 to buffered channel"
		bufferedChan <- "Message 2 to buffered channel" // 두 개의 메시지가 모두 버퍼에 저장되고, Goroutine 2는 블록되지 않는다
		fmt.Println("Goroutine 2: Data sent to buffered channel.")
	}()

	// Give some time for goroutines to execute
	time.Sleep(1 * time.Second) // time.Second는 1초, time.Sleep 함수는 인자로 받은 시간을 동안 현재 고루틴을 "수면 상태"로 만들어, 다른 작업을 할 수 없게 만든다.

	// Receiving from Unbuffered Channel
	fmt.Println("Main: Receiving data from unbuffered channel...")
	fmt.Println("Main: Received", <-unbufferedChan)

	// Receiving from Buffered Channel
	fmt.Println("Main: Receiving data from buffered channel...")
	fmt.Println("Main: Received", <-bufferedChan)
	fmt.Println("Main: Received", <-bufferedChan)
}
