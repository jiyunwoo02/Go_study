package main

import (
	"encoding/json" // 데이터를 JSON 형식으로 인코딩/디코딩하기 위한 패키지
	"net/http"      // HTTP 프로토콜을 이용한 서버/클라이언트 구현 위한 패키지
	"sort"          // 데이터를 정렬하는 함수를 제공하는 패키지
	"strconv"       // 문자열을 기본 데이터 타입으로 변환하는 기능 제공

	"github.com/gorilla/mux" // 강력한 URL 라우팅 기능을 제공하는 외부 패키지, 요청 URL과 관련된 핸들러를 매핑하는 데 사용
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

// 전역 변수 students : 학생의 ID를 키로 하여 Student 객체를 저장
var students map[int]Student

// 맵에 객체를 추가할 때 키 값으로 사용하는 것은 구조체의 어떤 필드든 될 수 있지만,
// 여기서는 고유 식별자의 역할을 하는 Id 필드가 자연스럽게 키로 사용된다!

var lastId int // 가장 최근에 추가된 마지막 학생의 ID를 추적

// 핸들러는 HTTP 요청을 받아서 그에 대응하는 작업을 수행하고, 결과를 클라이언트에게 돌려주는 역할

func MakeWebHandler() http.Handler { // 웹 서버의 요청을 처리할 라우터를 설정하고 반환
	mux := mux.NewRouter() // gorilla/mux 생성 [새로운 라우터 생성]
	// client가 특정 URL로 요청을 보낼 때, 해당 URL에 맞는 핸들러 함수를 찾아 실행해주는 것이 라우터의 역할

	// "/students" 경로로 "GET" 요청이 들어올 때 GetStudentListHandler 함수가 호출되도록 설정
	// 1. 학생 목록 조회하는 핸들러 함수와 경로 등록
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")

	//2. 특정 학생 정보 조회
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")

	// 3. 학생 데이터 추가
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")

	// 4. 학생 데이터 삭제
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	// 초기 학생 데이터 설정
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

/*
sort 패키지에 정의된 sort.Interface 인터페이스 구현
-> sort 패키지의 함수들을 사용하여 해당 슬라이스를 정렬할 수 있다.
  - sort.Interface는 Go 언어의 표준 라이브러리 sort 패키지에 정의된 인터페이스
  - 아래 세 메서드는 sort 패키지의 여러 정렬 함수들이 데이터를 정렬할 때 필요한 연산을 제공
  - 사용자가 자신의 타입에 대해 sort.Interface를 구현하면,
    -> 그 타입의 슬라이스에 대해 sort.Sort 등의 함수들을 사용할 수 있다
*/

func (s Students) Len() int { // 1) Len() int: 슬라이스의 길이를 반환
	return len(s)
}
func (s Students) Swap(i, j int) { // 2) Swap(i, j int): 슬라이스 내의 두 요소의 위치를 교환
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool { // 3) Less(i, j int) bool: 두 요소의 순서를 비교
	return s[i].Id < s[j].Id
	// i번째 요소가 j번째 요소보다 "작은지" 여부를 판단하여 정렬 순서를 결정
	// Id를 기준으로 학생들을 오름차순(ascend)으로 정렬
}

// 1. 저장된 모든 학생 정보를 JSON 형식으로 반환하는 핸들러
// /students 경로로 "GET" 요청이 들어왔을 때 호출되는 함수
func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
		sort.Sort(list) // ID를 기준으로 학생 목록을 정렬
		w.WriteHeader(http.StatusOK)
	}
	// students 맵에서 모든 학생 데이터를 Students 슬라이스에 추가한 다음, 이를 Id 순서대로 정렬
	// 응답의 상태 코드를 200 OK로 설정
	w.Header().Set("Content-Type", "application/json") // 응답의 콘텐츠 타입을 application/json으로 설정
	json.NewEncoder(w).Encode(list)                    // 정렬된 학생 목록을 JSON 형식으로 인코딩하여 응답 본문에 작성
}

// 2. 요청된 ID에 해당하는 학생 정보를 JSON 형식으로 반환
func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)               // mux.Vars() 호출해 인수 가져온다
	id, _ := strconv.Atoi(vars["id"]) // URL에서 id 파라미터를 정수로 변환
	student, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound) // 학생 데이터를 찾지 못한 경우 404 에러 반환
		return
	}

	w.WriteHeader(http.StatusOK)                       // 성공 상태 코드 설정
	w.Header().Set("Content-Type", "application/json") // 내용 타입 설정
	json.NewEncoder(w).Encode(student)                 // 학생 데이터를 JSON 형식으로 인코딩 후 응답
}

// 3. 새로운 학생 데이터를 받아서 시스템에 추가
func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student // 새로운 학생 정보를 저장할 Student 타입 변수 선언

	// 요청 본문에서 JSON 형식의 데이터를 읽어와 Student 구조체로 디코딩
	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // JSON 디코딩 실패 시, HTTP 400 (Bad Request) 상태 코드 반환
		return
	}
	lastId++            // 글로벌 lastId 변수를 증가시켜 새로운 학생에게 유니크한 ID 할당
	student.Id = lastId // 새로운 ID를 학생의 Id 필드에 설정

	students[lastId] = student // students 맵에 새 학생 정보를 추가

	w.WriteHeader(http.StatusCreated) // 데이터 추가 성공 시, HTTP 201 (Created) 상태 코드 반환
}

// 4. 주어진 ID에 해당하는 학생 데이터를 제거
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)               // URL 파라미터에서 데이터 추출
	id, _ := strconv.Atoi(vars["id"]) // URL 경로에 포함된 id 파라미터를 정수로 변환

	// students 맵에서 해당 ID를 키로 사용하여 학생 데이터 존재 여부 확인
	_, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound) // 해당 ID의 학생이 존재하지 않는 경우, HTTP 404 (Not Found) 반환
		return
	}
	// 해당 ID의 학생 데이터가 존재하는 경우, 맵에서 해당 학생 정보 삭제
	delete(students, id)
	w.WriteHeader(http.StatusOK) // 성공적으로 학생 데이터를 삭제한 경우, HTTP 200 (OK) 반환
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler()) // 3000번 포트에서 HTTP 서버 시작
	// MakeWebHandler에서 반환된 핸들러를 사용하여 요청 처리
	// http://localhost:3000/students로 요청을 보내면, 서버는 두 개의 미리 정의된 학생 데이터를 JSON 형식으로 응답
	// -> 이 데이터는 Id 기준으로 정렬된 상태로 반환
}
