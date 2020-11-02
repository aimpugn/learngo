package main

import (
	"fmt"
)

func main() {
	fmt.Println("====================addNumbers====================")
	a := addNumbers(2, 3)
	fmt.Println(a)
	a = addNumbers(4, 10)
	fmt.Println(a)
	a = addNumbers(100, -100)
	fmt.Println(a)

	fmt.Println("====================divAndRemainder====================")
	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)
	/*
		Go에서는 나중에 사용하지 않을 변수를 선언할 수 없다
		이때 언더스코어(_)를 사용하여 무시
	*/
	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)
	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)
	divAndRemainder(-1, 20)

	fmt.Println("====================doubleFail====================")
	// Go에서 모든 함수는 call by value, 따라서 함수 호출 시 복사된 변수가 함수로
	b := 1
	// arr1 := [2]int{2,4} << 이 경우 4 다음에 콤마는 불필요
	arr1 := [2]int{
		2,
		4, // << 이 경우 4 다음 콤마를 붙여야한다. 불이지 않으면 에러(syntax error: unexpected newline, expecting comma or })
	}
	s := "hello"
	doubleFail(b, arr1, s)
	fmt.Println("in main:", b, arr1, s)
}

// Go는 선택적 파라미터(optional parameter) 없다
// 함수의 파라미터에 이름을 붙여서 전달하는 name parameter(https://en.wikipedia.org/wiki/Named_parameter) 없다
func addNumbers(a int, b int) int { // return할 경우 ) 와 { 사이에 반환할 값의 타입 지정
	return a + b
}

/*
// function overloading 없다
func addNumbers(a int, b int, c int)  {  // ./main.go:22:6: addNumbers redeclared in this block
	fmt.Println(a + b + c)
}
*/

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Println("in doubleFail:", a, arr, s)
}
