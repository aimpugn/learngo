/*
첫 줄은 반드시 package 라인
main은 go 프로그램이 실행되는 특별한 패키지
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(r http.ResponseWriter, rq *http.Request) {
		name := rq.URL.Query().Get("name")
		r.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})

	http.ListenAndServe(":8080", nil)
}
