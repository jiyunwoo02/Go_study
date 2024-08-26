package main

import (
	"encoding/json" // 데이터를 JSON 형식으로 인코딩하고 디코딩하기 위한 패키지
	"net/http"      // HTTP 요청 및 응답을 처리하기 위한 Go의 표준 라이브러리
	"sort"          // 슬라이스의 데이터를 정렬하기 위한 패키지

	"github.com/gorilla/mux" // 강력한 URL 라우터 및 요청 디스패처를 제공하는 패키지, 요청 URL과 관련된 핸들러를 매핑하는 데 사용
	/* Go 공식 문서
	1. Package gorilla/mux implements a request router and dispatcher
	- for matching incoming requests to their respective handler.
	2. Gorilla is a web toolkit for the Go programming language
	- that provides useful, composable packages for writing HTTP-based applications.
	*/)

type Student struct { // 학생의 정보를 나타내는 구조체
	Id    int
	Name  string
	Age   int
	Score int
}

// 패키지 전역 변수 students : 학생 데이터를 저장하는 map
var students map[int]Student // Id를 키로, Student 데이터를 저장

var lastId int // 현재까지 추가된 마지막 학생의 ID를 추적

// 핸들러는 HTTP 요청을 받아서 그에 대응하는 작업을 수행하고, 결과를 클라이언트에게 돌려주는 역할
func MakeWebHandler() http.Handler { // 웹 서버의 핸들러를 생성하고 설정
	mux := mux.NewRouter() // gorilla/mux 생성 [새로운 라우터 생성]
	// client가 특정 URL로 요청을 보낼 때, 해당 URL에 맞는 핸들러 함수를 찾아 실행해주는 것이 라우터의 역할

	// "/students" 경로로 "GET" 요청이 들어올 때 GetStudentListHandler 함수가 호출되도록 설정
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")

	students = make(map[int]Student)

	// 임시 학생 데이터 두 개 생성해서 저장
	// GET 요청 받으면 -> GetStudentListHandler() 함수 호출
	// -> students 맵에 저장된 학생 데이터로 []Student 타입의 학생 목록 만든다
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}

	lastId = 2

	return mux
}

type Students []Student // Id로 정렬하는 인터페이스
// Students는 []Student라는 슬라이스 타입을 기반으로 만들어진 사용자 정의 타입
// []Student와 같은 데이터를 가지지만, 타입 이름을 통해 코드의 가독성을 높이고 특정 메서드를 추가할 수 있는 기능을 제공

// sort.Interface 인터페이스를 구현 -> Students 슬라이스 정렬
func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
	// Id를 기준으로 학생들을 오름차순으로 정렬
}

// 학생 정보를 가져와 JSON 포맷으로 변경하는 핸들러
// /students" 경로로 "GET" 요청이 들어왔을 때 호출되는 함수
func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0) // 학생 목록을 ID로 정렬 (맵은 키값에 따라 정렬되지 않고 저장된다)
	for _, student := range students {
		list = append(list, student)
		sort.Sort(list)
		w.WriteHeader(http.StatusOK)
	}
	// students 맵에서 모든 학생 데이터를 Students 슬라이스에 추가한 다음, 이를 Id 순서대로 정렬
	// 응답의 상태 코드를 200 OK로 설정
	w.Header().Set("Contemt-Type", "application/json") // 응답의 콘텐츠 타입을 application/json으로 설정
	json.NewEncoder(w).Encode(list)                    // 정렬된 학생 목록을 JSON 형식으로 인코딩하여 응답 본문에 작성
}

func main() {
	// 3000번 포트에서 입력 대기
	http.ListenAndServe(":3000", MakeWebHandler()) // 3000번 포트에서 HTTP 서버 시작, MakeWebHandler에서 반환된 핸들러를 사용하여 요청 처리
	// http://localhost:3000/students로 요청을 보내면, 서버는 두 개의 미리 정의된 학생 데이터를 JSON 형식으로 응답
	// -> 이 데이터는 Id 기준으로 정렬된 상태로 반환
}
