package main

import (
	"bufio"   // 문자열을 스캔하는 기능을 제공하는 패키지
	"fmt"     // 입출력 포매팅을 다루는 패키지
	"os"      // 운영 체제 기능, 특히 파일 입출력과 관련된 기능을 제공하는 패키지
	"strings" // 문자열 처리를 다루는 패키지
)

func main() {
	const input = `Now is the winter of our discontent,
	Made glorious summer by the sun of York.`

	// 문자열에서 스캐너 생성, bufio 스캐너를 사용하여 입력을 토큰화
	// bufio.NewScanner() 함수는 io.Reader 인터페이스를 인수로 받는다
	// -> string 타입 바로 사용할 수 X
	// -> strings.NewReader() 사용해서 문자열을 io.Reader 인터페이스 구현한 객체로 변환해서 넣어준다!
	scanner := bufio.NewScanner(strings.NewReader(input))

	// 스캔할 때 단어별로 분리 설정, 각 호출 시 스캐너는 입력에서 다음 단어를 읽음
	// 기본적으로 Scanner는 한 줄 단위로 토큰 읽는다
	// Split() 메서드 이용해 토큰 구분하는 함수 등록
	// bufio 패키지에서 제공하는 ScanWords() 메서드 이용해 단어 단위로 읽어온다
	scanner.Split(bufio.ScanWords)

	// 단어 수를 세기 위한 카운터 변수
	count := 0

	// 스캐너가 더 이상 단어를 읽을 수 없을 때까지 반복
	for scanner.Scan() {
		count++ // 스캔할 때마다 카운터 증가
		// Scan() 메서드가 false 반환하면 검색 중지
	}

	// 스캔 중에 발생할 수 있는 에러 처리
	if err := scanner.Err(); err != nil {
		// 에러가 있다면 표준 에러에 에러 메시지를 출력
		fmt.Fprintln(os.Stderr, "reading input:", err)
		// 더 읽을 수 없어서 검색 중단되면 Err() 메서드는 nil 반환
	}

	// 최종적으로 읽은 단어의 총 수를 출력 : 15
	fmt.Printf("%d\n", count)
}
