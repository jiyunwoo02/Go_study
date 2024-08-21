package main

func SendReport(r *Report, method sendType, receiver string) {
	switch method {
	case Email:
		// 이메일 전송
	case Fax:
		// 팩스 전송
	case PDF:
		// pdf 파일 생성
	case Printer:
		// 프린팅
	}
	// 전송 방식을 추가하려면 새로운 case를 만들어 구현을 추가해준다
	// -> 기존 SendReport() 함수 구현 변경하는 것
}

type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct{}

func (e *EmailSender) Send(r *Report) {
	// 이메일 전송
	// EmailSender는 ReportSender 인터페이스 구현한 객체
}

type FaxSender struct{}

func (f *FaxSender) Send(r *Report) {
	// 팩스 전송
	// FaxSender는 ReportSender 인터페이스 구현한 객체
}

// 새로운 전송 방식을 추가하려면?
// -> ReportSender를 구현한 새로운 객체를 추가해준다
// 새 기능을 추가했지만, 기존 구현을 변경하지 않아도 된다.
