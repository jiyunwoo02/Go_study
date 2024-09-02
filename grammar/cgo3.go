package main

/*
#include <stdio.h>
void hello() {
	printf("Hello from C");
}
*/
import "C"

func main() {
	// let's call it
	C.hello()
}


/*
1. import "C" 바로 위에 위치해야 한다
- 사이에 공백 있으면 오류 발생
PS > go run cgo3.go
# command-line-arguments
.\cgo3.go:15:2: could not determine kind of name for C.hello

2. 블록 주석 (/*) 말고 단일 주석 (//) 으로 하면 동일한 오류 발생
*/
