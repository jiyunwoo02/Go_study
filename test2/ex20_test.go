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
