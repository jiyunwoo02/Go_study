package main

type FinanceReport struct { // 회계 보고서
	report string
	// FinanceReport는 회계 보고서를 담당하는 객체
	// '회계 보고서'라는 책임을 진다.
}

func (r *FinanceReport) SendReport(email string) {
	// 회계 보고서 전송
	// FinanceReport가 보고서를 전송하는 책임까지 진다
	// => 책임이 2개가 되어 SRP 위배
}

type MarketingReport struct { // 마케팅 보고서
	report string
	// MarketingReport 객체 생성
	// FinanceReport의 SendReport 사용 불가능
}

func (r *MarketingReport) SendReport(email string) {
	// 마케팅 보고서 전송
	// 구현이 비슷한 메서드를 또 만들어야 한다..
	// 보고서 종류가 늘어날 때마다 SendReport()도 늘어나야 한다..
	// 또한 이메일이 아닌 다른 형태로 보내야 한다면 그동안 만든 메서드들을 모두 수정해야 한다.
}
