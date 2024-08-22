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
}
