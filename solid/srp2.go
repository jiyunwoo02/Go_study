package main

import "fmt"

// FinanceReport는 Report 인터페이스를 구현
// ReportSender는 Report 인터페이스를 이용하는 관계

type Report interface {
	Report() string // Report() 메서드 포함한 인터페이스
}

type FinanceReport struct {
	report string // 회계 보고서만을 담당하는 객체
}

func (r *FinanceReport) Report() string {
	return r.report // Report 인터페이스를 구현
}

type ReportSender struct { // 보고서 전송만 담당
}

func (s *ReportSender) SendReport(report Report) {
	fmt.Println("Sending report:", report.Report())
	// Report 인터페이스를 구현한 어떤 구조체도 받을 수 있다!
	// 향후 다른 보고서 나오더라도 그대로 ReportSender 이용 가능
}
