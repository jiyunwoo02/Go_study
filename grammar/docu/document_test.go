package doc_test

import "fmt"

// PrintDoc() 함수에 대한 예제입니다.
func ExamplePrintDoc() {
	fmt.Println("This is an package level example")
}

// TextDoc의 PrintDoc() 메서드에 대한 예제입니다.
func ExampleTestDoc_PrintDoc() {
	fmt.Println("This is PrintDoc() example")
}

// TextDoc에 대한 예제입니다.
func ExampleTextDoc_lines() {
	fmt.Println("This is lines() example")
}
