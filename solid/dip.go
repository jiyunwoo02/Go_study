package main

import "fmt"

type Mail struct {
	alarm Alarm
	// 메일 객체는 알람 객체 소유
	// Mail이라는 상위 모듈이
	// -> Alarm이라는 하위 모듈에 의존
}
type Alarm struct{}

func (a *Alarm) Alarm() {
	// 알람을 울리는 메서드
	fmt.Println("Alarm is ringing!")
}
func (m *Mail) onRecv() {
	m.alarm.Alarm()
	// Mail은 Alarm의 메서드 Alarm()을 직접 호출
	// Mail은 Alarm의 구체적인 구현에 의존
	// 메일 수신 시 OnRecv() 메서드에서
	// -> 소유한 알람 객체 사용해 알람 울린다
}
