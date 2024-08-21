package main

import "time"

type Report interface { // 총 4개의 메서드 포함
	Report() string
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) { // Report() 메서드만 사용
	send(r.Report())
	// 인터페이스 이용자에게 불필요한 메서드들을
	// 인터페이스가 포함하고 있다.
}

type Report interface {
	Report() string
}
type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}
