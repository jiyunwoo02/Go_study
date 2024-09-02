package main

/*
#include <math.h>
*/
import "C"
import (
	"fmt"
)

// CMathSin 함수는 입력된 각도(라디안)에 대한 사인 값을 계산
func CMathSin(radians float64) float64 {
	return float64(C.sin(C.double(radians)))
}

// CMathCos 함수는 입력된 각도(라디안)에 대한 코사인 값을 계산
func CMathCos(radians float64) float64 {
	return float64(C.cos(C.double(radians)))
}

// CMathSqrt 함수는 입력된 값의 제곱근을 계산
func CMathSqrt(value float64) float64 {
	return float64(C.sqrt(C.double(value)))
}

func main() {
	fmt.Printf("Sin(π/2) = %v\n", CMathSin(3.14159265/2))
	fmt.Printf("Cos(π) = %v\n", CMathCos(3.14159265))
	fmt.Printf("Sqrt(16) = %v\n", CMathSqrt(16))
}
