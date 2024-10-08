/* < 1. Tuple assignment 이 실행되는 세부 순서 Ex) x, y := 0, 1 ? >

- tuple assignment: 여러 개의 값을 동시에 할당하는 방법, 여러 변수에 여러 값을 한 번에 할당할 수 있는 구문

주로 함수가 여러 값을 반환할 때나, 여러 변수에 값을 동시에 할당할 때 사용한다
- Go 코드의 가독성 높인다, 여러 값을 효율적으로 처리하는 데 유용한 기능

Go에서는 "튜플(Tuple)"이라는 용어를 직접 사용하지는 않지만, 이 개념에서의 "튜플"은 사실 여러 개의 값을 한 번에 처리하는 것을 의미
- Go에서 "튜플"은 일종의 다중값(multiple values) 을 뜻한다
+ Go에서는 반환값의 자료형이 반드시 동일할 필요는 없다
(예를 들어, 두 개의 반환값을 반환하는 함수에서 하나는 int, 하나는 string 가능)

+ Python의 "튜플": 여러 개의 항목을 담을 수 있는 컬렉션으로, 항목들이 순서대로 저장되며 변경할 수 없는(immutable) 특징을 가진다

[세부 순서]
1) 오른쪽 값들의 평가
: 먼저 할당 연산자(=)의 오른쪽에 있는 모든 값들이 순서대로 평가된다.
: 각 값이 독립적으로 계산되고, 결과 값들이 임시 저장소에 저장된다. <- 이 때도 메모리가 할당된다.

2) 왼쪽 변수들의 평가
: 그 다음으로, 할당 연산자의 왼쪽에 있는 변수들이 순서대로 평가된다.
: 만약 왼쪽에 기존의 변수가 있다면 그 변수의 메모리 위치가 확인되고, 만약 새로운 변수가 선언된 경우에는 메모리가 할당된다.
- 변수는 선언될 때 메모리 주소를 할당 받는다

3) 할당
: 마지막으로, 오른쪽에서 평가된 값들이 왼쪽의 변수들에 동시에 할당된다.
: 이 단계에서는 모든 변수가 이전에 평가된 오른쪽 값들을 받는다.
- 이 동시 할당의 주요 특징은 '모든 값이 동시에 변경된다는 것' -> Go에서 tuple assignment의 중요 특성

임시 저장소에 저장된 값 자체가 좌측 변수들이 가리키는 메모리 위치로 복사된다!
- 임시 저장소의 값 자체가 좌측 변수에 할당된 메모리로 옮겨지며, 임시 저장소의 메모리 주소는 대입되지 않는다.
-- 즉, 좌측 변수들은 임시 저장소의 메모리 주소가 아닌, 임시 저장소에 저장된 값을 자신의 메모리 위치로 복사받는다.

주의할 점) 임시 값의 사용
: Go는 할당 과정에서 임시 저장소를 사용하여 오른쪽 값들을 임시로 저장한다.
: 이는 할당 과정에서 값이 서로 덮어씌워지는 것을 방지하기 위함이다.

:= 연산자는 변수를 선언하면서 동시에 초기화할 때 사용
*/

package main

import (
	"fmt"
)

func swap(x, y int) (int, int) {
	// x와 y의 메모리 주소 출력: swap 함수 내부에서 할당된 x와 y의 메모리 위치
	fmt.Printf("swap 함수의 x의 메모리 주소: %p\n", &x)
	fmt.Printf("swap 함수의 y의 메모리 주소: %p\n\n", &y)

	// 두 개의 값을 반환, main 함수로 반환될 값들을 전달하기 위해 임시 저장소에 저장
	return y, x

	// x와 y의 메모리 공간은 반환 후에도 원래의 값을 (1과 2) 가지고 있으며,
	// 반환된 후의 임시 저장소에 있는 값이 a와 b에 복사된다
}

func main() {
	fmt.Println("Tuple assignment에 대해 알아보자\n")

	// 1. 다중 값 반환 처리
	a, b := swap(1, 2)
	// (1) swap(1, 2)가 호출되면, 1과 2라는 값이 swap 함수의 파라미터 x와 y로 전달
	// (2) 1과 2는 각각 x와 y라는 새로운 변수에 값이 복사됨
	// (3) x와 y는 함수 호출 시점에 메모리 공간을 할당받고, 각각의 메모리 공간에 1과 2가 저장
	// (4) y는 2로, x는 1로 평가되어 임시 저장소에 저장됨
	fmt.Printf("swap 함수의 결과 -> a : %d, b : %d\n", a, b) // 출력: 2 1
	// (5) 좌측의 a와 b는 새로 선언되는 변수이므로, 각각 메모리 공간이 새로 할당됨
	// (6) 임시 저장소에 있는 값들이 좌측 변수 a와 b에 각각 복사됨

	fmt.Printf("a의 메모리 주소: %p\n", &a)   // a의 메모리 주소 출력
	fmt.Printf("b의 메모리 주소: %p\n\n", &b) // b의 메모리 주소 출력
	// -> swap 함수 내에서 사용되는 x와 y의 메모리 주소는 main 함수의 a와 b와는 별개의 메모리 주소!

	// 2. 여러 변수에 한꺼번에 값 할당
	x1, y1, z1 := 1, 2, 3 // 변수의 타입은 오른쪽의 값들에 따라 자동으로 유추 (int)
	// x1, y1, z1라는 세 개의 변수가 동시에 선언되고 -> 메모리 할당받고 -> 각각 1, 2, 3의 값이 할당
	fmt.Printf("x1: %d, y1 : %d, z1 : %d\n", x1, y1, z1) // 출력: 1 2 3
	fmt.Printf("x1의 메모리 주소: %p\n", &x1)

	// 3. 기존 변수와 새 변수에 값 할당 및 업데이트
	x2 := 1
	y2 := 2
	// 1) x2, y2라는 두 개의 변수가 각각 1, 2로 초기화
	// 2) x2와 y2는 각각의 메모리 공간에 1과 2의 값을 저장
	x2, z2 := 3, 4
	// x2는 기존 변수, z2는 새 변수 -> 이 경우, x2는 기존에 선언된 변수로 값을 업데이트하고, z2는 새 변수로 할당
	// 3) x2는 이미 선언된 변수이므로, 새로운 값을 할당받아 3으로 업데이트
	// 4) z2는 새롭게 선언된 변수로, 4의 값을 할당받음
	fmt.Printf("x2: %d, y2 : %d, z2 : %d\n", x2, y2, z2) // 출력: 3 2 4
}

/* swap

1. 함수 호출 및 다중 반환
먼저, `swap` 함수가 호출:

- swap(1, 2) 호출 시:
  - `x`는 `1`로 초기화되고, `y`는 `2`로 초기화
  - `return y, x`에 의해 함수는 두 값을 반환, 이때 반환 순서는 `(2, 1)`

이 반환된 두 값은 임시적으로 메모리에 저장된다

2. 튜플 할당
`a, b := swap(1, 2)`에서 일어나는 동작:

1) 새 변수 생성
   - `a`와 `b`는 새롭게 선언된 변수
   - Go에서는 `:=` 연산자를 사용할 때, 새로운 변수에 메모리 공간이 할당.
   - 이 경우 `a`와 `b`가 각각 새로운 메모리 주소를 갖게 된다

2) 값 복사
   - `swap(1, 2)`가 반환한 `(2, 1)`의 값들이 각각 `a`와 `b`에 복사
   - 즉, `a`는 `2`, `b`는 `1`로 설정
   - 이 과정에서 임시 저장된 `(2, 1)`의 값은 `a`와 `b`의 메모리 주소에 복사

3) 할당 완료
   - 이제 `a`는 `2`, `b`는 `1`을 갖고, 각각의 메모리 주소에 이 값들이 저장되어 있다.

3. 메모리에서의 동작

1) 함수 호출 시
   - `swap(1, 2)`가 호출되면, `x`와 `y`가 각각 메모리의 특정 위치에 `1`과 `2`로 저장된다

2) 함수 반환 시
   - `return y, x`가 실행되면, `2`와 `1`이 임시 메모리 위치에 저장된다

3) 튜플 할당 시
   - 새로운 메모리 공간이 `a`와 `b`에 할당된다
   - 임시 메모리에 저장된 값 `2`는 `a`에, `1`은 `b`에 복사된다

=> 정리
1. `swap(1, 2)` 함수 호출로 인해 `(2, 1)`이 반환
2. 반환된 값 `(2, 1)`이 각각 `a`와 `b`에 할당
3. `a`는 `2`, `b`는 `1`로 설정

*/

/* 추가 의문

1. 좌측 변수와 우측 값의 개수가 일치하지 않으면?
- 좌측 변수와 우측 값의 개수가 일치하지 않으면 컴파일 에러가 발생
- 모든 좌측 변수에 대해 하나의 우측 값이 필요하며, 그렇지 않으면 Go 컴파일러가 오류를 보고한다.
- 필요 없는 값은 _를 사용해 무시할 수 있다.
- a, b, _ := values() // values()는 3개의 정수(1,2,3) 리턴, 세 번째 값은 무시

2. 좌측과 우측에 모두 변수 사용 가능여부
- 이는 특히 값의 교환이나 연산 후의 재할당과 같은 상황에서 유용
- x := 5, x = x + 2

3. 메모리 주소 취득 시점
1) 매개변수는 함수가 호출되면서 메모리 공간이 할당
2) 함수의 인자로 전달된 값들이 해당 함수의 매개변수 x와 y에 복사되는 시점에 x와 y는 새로운 메모리 공간을 할당받는다
*/
