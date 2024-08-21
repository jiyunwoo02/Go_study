package main

type T interface {
	Something()
	// Something 메서드 포함한 인터페이스
}
type S struct{}

func (s *S) Something() {
	// T 인터페이스 구현
}

type U struct{}

func (u *U) Something() {
	// T 인터페이스 구현
}
func q(t T) {}

var y = &S{} // S 타입 y
var u = &U{} // U 타입 u
// q(y)와 q(u) 둘 다 잘 동작해야 한다
// S와 U가 T의 하위 타입이기 때문에 상위 타입인 T를 인수로 받는 함수에
// -> 인스턴스를 넣어도 잘 동작해야 한다
