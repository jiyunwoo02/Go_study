package main

import (
	"fmt"
	"sync"
)

// 숫자 리스트 합계를 계산하는 인터페이스 정의
type SumCalculator interface {
	Add(numbers []int) // 숫자 리스트를 받아 합계에 추가하는 메서드
	Result() int       // 현재 합계를 반환하는 메서드
}

// 숫자 리스트 합계를 계산하는 구조체 정의
type SimpleSumCalculator struct {
	sum int // 합계를 저장할 필드
	//mu  sync.Mutex // 동시 접근을 방지할 뮤텍스
}

// Add 메서드: 숫자 리스트를 받아 합계에 추가
// s *SimpleSumCalculator는 리시버.
// s는 *SimpleSumCalculator 타입
// *SimpleSumCalculator는 메서드가 SimpleSumCalculator 구조체의 포인터를 통해 호출된다는 것
// -- 포인터 리시버: 포인터를 사용하면 구조체의 필드를 직접 수정 가능
func (s *SimpleSumCalculator) Add(numbers []int) {
	//s.mu.Lock()         // 뮤텍스 잠금
	//defer s.mu.Unlock() // 메서드가 끝날 때 뮤텍스 잠금 해제
	for _, number := range numbers {
		s.sum += number // 각 숫자를 합계에 추가
	}
}

// Result 메서드: 현재 합계 반환
func (s *SimpleSumCalculator) Result() int {
	//s.mu.Lock()         // 뮤텍스 잠금
	//defer s.mu.Unlock() // 메서드가 끝날 때 뮤텍스 잠금 해제
	return s.sum // 현재 합계를 반환
}

func main() {
	// 고루틴 동기화를 위한 WaitGroup: 모든 고루틴의 작업이 완료될 때까지 대기
	var wg sync.WaitGroup
	calculator := &SimpleSumCalculator{} // SimpleSumCalculator 구조체의 인스턴스 생성

	// 숫자 리스트를 여러 부분으로 나누어 각 고루틴에서 처리
	parts := [][]int{ // 묶음의 묶음
		{1, 2, 3, 4, 5},      // 부분 1: 묶음
		{6, 7, 8, 9, 10},     // 부분 2
		{11, 12, 13, 14, 15}, // 부분 3
		{16, 17, 18, 19, 20}, // 부분 4
	}

	// 채널을 통해 각 고루틴의 작업 완료 신호를 수신
	doneChannel := make(chan bool)

	// 각 부분을 처리할 고루틴 생성
	for _, part := range parts { // i는 0부터 시작한다
		wg.Add(1)             // WaitGroup에 작업 추가
		go func(nums []int) { // 고루틴은 순서가 뒤죽박죽, 스케쥴러가 적당한 타이밍에 실행시켜주는데, 실행 순서 보장할 수 없다
			defer wg.Done()      // 고루틴 작업이 끝나면 WaitGroup에서 작업 완료
			calculator.Add(nums) // 숫자 리스트를 합계에 추가
			doneChannel <- true  // 작업 완료 신호 전송
		}(part) // 현재 숫자 리스트를 고루틴에 전달
	}

	// 모든 고루틴의 작업이 완료될 때까지 대기
	go func() {
		wg.Wait()          // 모든 고루틴의 작업이 완료될 때까지 대기
		close(doneChannel) // 채널 닫기 (모든 작업 완료 신호 수신)
	}()

	// 채널을 통해 모든 고루틴의 작업 완료 신호 수신
	for range doneChannel {
		// 신호 수신 후 처리 (현재는 단순히 작업 완료 확인)
		fmt.Println("확인")
	}

	// 최종 합계 출력 => "최종 합계: 210"
	fmt.Printf("최종 합계: %d\n", calculator.Result())
}
