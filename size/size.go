package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

func main() {
	// 시스템 아키텍처 출력
	fmt.Printf("System Architecture: %s\n", runtime.GOARCH)

	// 타입 크기 출력
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("Size of string: %d bytes\n", unsafe.Sizeof(""))
	fmt.Printf("Size of sync.Mutex: %d bytes\n", unsafe.Sizeof(sync.Mutex{}))
}
