package main

import (
	"fmt"
	"sync"
	"time"
)

// ex1.go 의 각 고루틴을 별개의 함수로 분리

func func1(wg *sync.WaitGroup) {
	defer wg.Done()             // Done() 호출
	time.Sleep(1 * time.Second) // 잠시 대기
	fmt.Println("고루틴 1 작업 완료")
}

func func2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("고루틴 2 작업 완료")
}

func func3() {
	// Done() 호출 X
	time.Sleep(2 * time.Second)
	fmt.Println("고루틴 3 작업 완료")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // Done() 호출이 2번 있어야 Wait()이 종료됨

	// 고루틴 1
	go func1(&wg)

	// 고루틴 2
	go func2(&wg)

	// 고루틴 3
	go func3()

	wg.Done() // 의도적으로 추가 Done 호출 (3번째 Done)

	// Wait()는 Done()이 2번 호출될 때까지 기다림
	wg.Wait()

	fmt.Println("작업이 완료되었습니다.")
}

/*
위 코드를 실행하면 프로그램이 오류 없이 정상적으로 종료되는 경우도 존재하지만,

예) PS C:\Users\jiyun\OneDrive\바탕 화면\goproject\example> go run ex2.go
고루틴 2 작업 완료
작업이 완료되었습니다.

아래와 같이 패닉이 발생하는 경우도 존재한다. 해당 문구를 살펴보자.

예) PS C:\Users\jiyun\OneDrive\바탕 화면\goproject\example> go run ex2.go
고루틴 2 작업 완료
고루틴 1 작업 완료
작업이 완료되었습니다.
panic: sync: negative WaitGroup counter

goroutine 6 [running]:
sync.(*WaitGroup).Add(0xc00008e000?, 0x0?)
        C:/Program Files/Go/src/sync/waitgroup.go:62 +0xd8
sync.(*WaitGroup).Done(...)
        C:/Program Files/Go/src/sync/waitgroup.go:87
main.func1(0x0?)
        C:/Users/jiyun/OneDrive/바탕 화면/goproject/example/ex2.go:15 +0x9b
created by main.main in goroutine 1
        C:/Users/jiyun/OneDrive/바탕 화면/goproject/example/ex2.go:35 +0x69
exit status 2

(1) panic : 프로그램이 런타임에서 오류를 만나 실행을 중단하고, 패닉 상태로 진입

(2) sync: negative WaitGroup counter: sync.WaitGroup의 내부 카운터 값이 음수가 되었다는 오류
- Done() 메서드가 Add() 메서드로 추가된 작업보다 더 많이 호출되었을 때 발생

(3) goroutine 6 [running]: 패닉이 발생한 시점에 실행 중이던 고루틴의 정보
- 이 고루틴이 Done() 메서드를 호출하면서 WaitGroup의 카운터를 감소시키려고 했을 때 오류가 발생

(4) sync.(*WaitGroup).Add(0xc00008e000?, 0x0?): sync.WaitGroup의 Add 메서드가 호출된 부분의 코드
- Add()는 WaitGroup의 카운터를 변경하는 함수이며, 이 줄에서 음수 카운터가 감지되어 패닉 발생

(5) sync.(*WaitGroup).Done(...): Done() 메서드가 호출된 부분
- Done()은 Add(-1)을 내부적으로 호출해 WaitGroup의 카운터를 감소

(6) main.func1(0x0?): 패닉이 발생한 고루틴에서 실행 중이던 함수
- func1에서 wg.Done()이 호출된 것
- ex2.go:15는 이 함수가 정의된 파일과 줄 번호

(7) created by main.main in goroutine 1: main 함수에서 고루틴이 생성된 것
- 이 고루틴은 main.main에서 시작되었으며, ex2.go 파일의 35번째 줄에서 생성됨
- package main에 정의된, Go 프로그램의 진입점이자 실행되는 최초의 함수인 main 함수

(8) exit status 2: 프로그램이 비정상 종료되었음을 나타내는 종료 상태 코드
*/

/*
1. 패닉은 언제 발생하는가?
: 패닉은 sync.WaitGroup의 Add() 메서드에서 발생한다.

구체적으로, Add(-1)이 호출될 때 WaitGroup의 내부 카운터가 음수로 감소하면,
이 상황을 감지한 Go 런타임이 패닉을 일으킨다.
- WaitGroup의 Add(delta int) 메서드는 delta만큼 내부 카운터를 증가시키거나 감소시킨다.
- Done() 메서드는 내부적으로 Add(-1)을 호출하여 카운터를 감소시킨다.
- 카운터는 고루틴들이 모두 작업을 완료하고 Done()을 호출할 때마다 감소한다.

문제: 만약 Done()이 예상보다 더 많이 호출되거나, Add()에서 양수로 카운터를 증가시키는 호출이 적절히 이루어지지 않으면, 카운터가 음수로 내려간다.
결과: 카운터가 음수로 내려가는 순간, WaitGroup은 이를 비정상적인 상태로 인식하고, 이 시점에서 패닉이 발생한다.
=> 그래서 Add()를 호출할 때, 내부적으로 delta 값을 이용해 카운터를 조작하는데, 이 과정에서 카운터가 음수가 되면 패닉을 일으키는 것이다.

*/

/*
추가적으로
1) Add()와 Done()
- 모든 고루틴이 자신의 작업을 완료한 후 Done()을 호출하는 것은 필수가 아니다.
: 3개의 고루틴이 있는데, Add(2)를 실행했으면 Done()이 2번만 이루어지면 된다.

2) 고루틴6은 어디서?
- 고루틴 번호는 Go 런타임에서 관리하는 내부적인 요소이며,
이 번호는 사용자 코드에서 생성한 고루틴 수와 일치하지 않을 수 있다.
- 프로그램 내에서 네 개의 고루틴(메인 고루틴 포함)이 생성되었지만,
goroutine 6이 나온 것은 Go 런타임이 내부적으로 관리하는
다른 고루틴들(예: 가비지 컬렉터, 스케줄러 등)이 먼저 생성되었을 수 있기 때문이다!
: goroutine 6 [running]이라는 메시지는 Go 런타임이 현재 실행 중인 고루틴을 표시하는 것이며,
이는 프로그램 내의 전체 고루틴 수와는 무관하다.

3) wg *sync.WaitGroup 와 &wg
- wg라는 매개변수가 *sync.WaitGroup 타입
- *sync.WaitGroup은 sync.WaitGroup 구조체를 가리키는 포인터
- 포인터를 사용하면 함수 내부에서 매개변수로 받은 구조체의 원래 값을 직접 수정할 수 있다.
- 포인터를 통해 전달된 변수는 함수 내부에서 변경되더라도 함수 밖에서 그 변경 사항이 반영된다.

- &wg는 wg라는 변수를 가리키는 주소(메모리 위치)를 인자로 전달
- & 연산자는 해당 변수의 주소를 반환하는 연산자
- go func2(&wg)를 호출 -> func2 함수가 wg의 메모리 주소를 통해 직접 wg 변수에 접근할 수 있다.
- 이로 인해 함수 func2는 wg의 원래 값을 수정할 수 있게 된다.

4) Add()와 Wait()가 동시에 호출되면 panic
: Add는 새로운 작업을 추가하는 과정에서, 작업 카운터를 조작하고,
Wait는 모든 작업이 완료되기를 기다리며 카운터를 검사한다.
이 두 작업이 동시에 발생하면, 작업 카운터의 값이 예상치 못한 방식으로 변경될 수 있으므로,
코드에서는 이를 엄격하게 관리하고 있다.

5) 만약 add가 done보다 많으면?
: Add의 호출 수가 Done보다 많다면, 이 카운터는 결코 0에 도달하지 않으므로 Wait는 영원히 반환되지 않는다.
- 무한 대기 or 프로그램 데드락 감지 or 오류.경고 메시지
*/
