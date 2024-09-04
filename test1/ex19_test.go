package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare3(t *testing.T) {
	assert := assert.New(t) // 테스트 객체 생성
	// assert 객체는 테스트 코드를 쉽게 만들 수 있는 메서드 포함
	// assert 객체에서 제공하는 Equal() 메서드 사용
	// 그 외 NotEqual(), Nil(), NotNil()
	assert.Equal(81, Square(9), "Square(9) should be 81")
}

func TestSquare4(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, Square(3), "Square(3) should be 9")
	// assert.Equal(4, Square(3), "Square(3) should be 9")
}

/*
--- FAIL: TestSquare4 (0.00s)
    ex19_test.go:19:
                Error Trace:    C:/Users/jiyun/OneDrive/바탕 화면/goproject/test1/ex19_test.go:19
                Error:          Not equal:
                                expected: 4
                                actual  : 9
                Test:           TestSquare4
                Messages:       Square(3) should be 9
FAIL
exit status 1
FAIL    goproject/test1 0.821s
*/
