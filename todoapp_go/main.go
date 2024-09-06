package main

import (
	"html/template"
	"net/http"
	"sync"
)

// Todo 항목을 정의하는 구조체
type Todo struct {
	Title string
	Done  bool
}

// Todo 목록을 관리하는 구조체
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var todoList []Todo // 할 일 목록을 저장하는 슬라이스
var mu sync.Mutex   // 동시성 관리를 위한 뮤텍스

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html")) // HTML 템플릿 파싱

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     todoList,
		}
		mu.Unlock()
		tmpl.Execute(w, data) // 템플릿 렌더링
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		if title != "" {
			mu.Lock()
			todoList = append(todoList, Todo{Title: title})
			mu.Unlock()
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}
