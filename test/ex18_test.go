package main

import "testing"

// 테스트 코드: 9의 제곱값이 81임을 테스트
// 터미널에서 go test 입력
// 혹은 VScode에서 자동으로 나오는 run package test 클릭해 실행
func TestSquare(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) shoud be 81 but returns %d", rst)
		// t.Errorf() 메서드에 테스트 실패 시 실패를 알리고 실패 메시지 넣을 수 있다
		// testing.T 객체의 Error()와 Fail() 메서드 이용해 테스트 실패 알린다
		// Error()는 테스트 실패하면 모든 테스트 중단
		// Fail()은 테스트 실패해도 다른 테스트 계속 진행
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) should be 9 but returns %d", rst)
	}
	// 두 테스트 다 테스트 성공하면 두 테스트 한 번에 실행해도 하나의 pass ok 뜬다
	// 테스트 하나씩 실행하려면 go test -run Square1
}
