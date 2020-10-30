/*
##### 구조 #####
- cmd/서비스/main.go
- package는 서비스명이 아닌 main
*/
package main

import "fmt"

func main() {
	// 1. 실행: go run main.go
	fmt.Println("Hello, World!")
	// 결과: Hello, World!

	// 2. 빌드
	// 2020-10-28  오후 11:23         2,143,744 main.exe
	// 2020-10-28  오후 11:23               261 main.go
	// C:\Users\daybreak\GolandProjects\stsa\cmd\a_tutorial>main.exe
	// Hello, World!
}
