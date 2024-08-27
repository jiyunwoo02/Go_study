package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t) // 테스트 객체 생성

	res := httptest.NewRecorder()
	// HTTP 응답을 기록할 ResponseRecorder 생성, "테스트에서 서버의 응답을 수집"하는 데 사용

	req := httptest.NewRequest("GET", "/students", nil)
	// "/students" 경로에 대한 GET 요청을 생성, 요청 본문은 없다 (nil)
	// httptest.NewRequest는 "테스트 목적으로 사용되는 HTTP 요청을 생성"

	mux := MakeWebHandler() // 실제 웹 핸들러 생성

	mux.ServeHTTP(res, req)
	// ServeHTTP는 요청(req)을 처리하고, 응답을 res에 기록
	// "/students" 경로에 대한 GET 요청을 핸들러가 처리

	assert.Equal(http.StatusOK, res.Code)
	// 응답 코드가 200 OK인지 확인

	var list []Student
	// 응답 본문을 JSON으로 디코딩할 Student 슬라이스

	err := json.NewDecoder(res.Body).Decode(&list)
	// 응답 본문을 JSON으로 디코딩하여 list라는 Student 슬라이스에 저장
	// res.Body는 HTTP 응답의 본문

	assert.Nil(err)
	// JSON 디코딩 중 오류가 발생하지 않았는지 확인

	assert.Equal(2, len(list))
	// 디코딩된 Student 슬라이스의 길이가 2인지 확인

	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}
