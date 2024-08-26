package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 테스트 코드 작성
func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(3) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}
func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibonacci2(3) should be 2")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}

// 벤치마크 코드
// 피보나치 수열 - 재귀 호출과 반복문 성능 측정 비교
// 피보나치 수열: 첫째/둘째 항은 1, 그 뒤의 항은 바로 앞 두 항의 합
func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// fibonacci1(20)을 b.N만큼 반복
		// Go는 N값 적절히 증가시키며 테스트해 성능 측정
		fibonacci1(20)
	}
}
func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}

/*
func BenchmarkXxxx(b *testing.B)
: 벤치마크 코드를 작성할 때, testing 패키지의 *testing.B 타입을 테스트 함수의 매개변수로 사용
- B는 testing 패키지에서 정의된 구조체 타입
- 벤치마크 테스트를 실행하고, 성능 데이터를 수집하고, 벤치마크에 관련된 여러 가지 메서드를 제공
- 벤치마크를 여러 번 반복하여 실행하여 평균 시간을 측정
- b.N 값을 통해 벤치마크 반복 횟수를 제어
- 벤치마크 테스트의 출력을 기록하고 관리
*/

/*
3. ex20_test.go
두 벤치마크 코드에서는 for문을 사용해  fibonacci(20) 함수를 b.N 번 반복해서 호출
- b.N은 Go 벤치마크 테스트에서 벤치마크 함수가 반복되는 횟수를 나타낸다. (N은 Number 약자)
- Go 벤치마크 프레임워크가 함수의 실행 시간을 측정하고, 일정한 시간 내에 실행될 수 있는 반복 횟수를 조절!
- b.N이 증가하면서, fibonacci 함수가 반복적으로 실행되고, 이를 통해 해당 함수의 성능을 측정
*/
