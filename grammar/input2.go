package main

import (
	"fmt"
	"os"
)

func main() {
	// os 패키지의 Create() 함수 이용해서 파일 생성
	// Create() 함수는 파일 이미 있으면 삭제, 없으면 생성
	f, err := os.Create("ouptut.txt")

	// 파일 생성 과정에서 오류 발생 시 에러 반환
	if err != nil {
		// 표준 오류로 출력
		fmt.Errorf("Create : %v\n", err)
		return
		// 파일 생성 성공할 경우 파일 핸들 객체인 *File 반환
		// *File 타입은 io.Writer를 구현하고 있기 때문에 Fprint() 함수의 인수로 사용할 수 있다
	}

	defer f.Close()

	const name, age = "kim", 22

	// 원하는 문자열 형태를 파일에 쓴다
	// Fprint() 함수 이용해 원하는 포맷으로 io.Writer 객체에 문자열 쓸 수 있다
	n, err := fmt.Fprint(f, name, " is ", age, " years old.\n")

	if err != nil {
		fmt.Errorf("Fprint: %v\n", err)
	}

	// ouptut.txt 파일 생성되고, kim is 22 years old. 적혀있다
	fmt.Print(n, " bytes written\n") // 21 bytes written
}
