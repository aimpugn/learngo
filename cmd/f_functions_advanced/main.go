package main

import (
	"fmt"
)

func addOne(a int) int {
	return a + 1
}

func assignLocalFunc() {
	/*
		함수를 로컬 변수에 할당
		()를 쓰지 않음으로써 함수 호출하는 것이 아니라, 함수를 참조(reference)
	*/
	myAddOne := addOne
	fmt.Println(addOne(1))   // 2
	fmt.Println(myAddOne(1)) // 2
}

func assignAnonymousFunc() {
	myAddOne := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddOne(1))
}

/*
func declareNestedFunc()  {
	// Unresolved reference 'nestedInFunc'
	func nestedFunc(a int) int {
		return a + 1
	}
}
*/

func addTwo(a int) int {
	return a + 2
}

// f func(int) int: int 하나를 매개변수로 받고, int 하나를 반환하는 어떤 함수든 전달 가능
func useFuncParameter(a int, f func(int) int) {
	fmt.Println(f(a))
}

func useNestedAnonymousFuncVariable() {
	b := 2
	myAddOne := func(a int) int { // 반환 타입 int 지정했는데 return 안하면 에러 발생(missing return at end of function)
		return a + b
	}
	fmt.Println(myAddOne(1)) // 3
}

func assignLocalVariableInNestedAnonymousFunc() {
	b := 2
	myAddOne := func(a int) { //
		b = a + b
	}
	myAddOne(1)
	fmt.Println(b) // 3
	/*
		클로저(closures)
		로컬 함수는 그 함수가 선언된 환경(environment)에 선언된 변수에 접근(access)할 수 있다
	*/
}

func makeAdderFunc(b int) func(int) int {
	return func(a int) int { // 클로저(closures)
		return a + b
	}
}

func makeDoublerFunc(f func(int) int) func(int) int {
	return func(a int) int { // 클로저(closures)
		b := f(a)
		return b * 2
	}
}

func main() {
	fmt.Println("==============assignLocalFunc==============")
	assignLocalFunc()
	fmt.Println("==============assignAnonymousFunc==============")
	assignAnonymousFunc()
	fmt.Println("==============useFuncParameter==============")
	useFuncParameter(1, addOne) // 2
	useFuncParameter(1, addTwo) // 3
	// useFuncParameter(1, assignAnonymousFunc)  // Cannot use 'assignAnonymousFunc' (type func()) as type func(int) int
	fmt.Println("==============useNestedAnonymousFuncVariable==============")
	useNestedAnonymousFuncVariable()
	fmt.Println("==============assignLocalVariableInNestedAnonymousFunc==============")
	assignLocalVariableInNestedAnonymousFunc() // 클로저(closures)
	fmt.Println("==============makeAdderFunc==============")
	/*
		원래 함수에서 전달된 변수는 함수가 종료되면 사라진다
		하지만 클로저에서는 makeAdderFunc로 반환된 함수 addOne와 addTwo는 전달된 파라미터 반환된 클로저의 b=1, b=2 값을 유지한다
	*/
	addOne := makeAdderFunc(1)
	fmt.Println(addOne(1)) // 2
	addTwo := makeAdderFunc(2)
	fmt.Println(addTwo(1)) // 3
	fmt.Println("==============makeDoublerFunc==============")
	doubleAddOne := makeDoublerFunc(addOne)
	fmt.Println(addOne(1))       // 2
	fmt.Println(doubleAddOne(1)) // 4

	/*
		클로저의 사용례
		http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		> 어떤 URI 경로가 호출될 때마다 특정 코드 실행하고 싶을 때 사용 가능
			> http.HandleFunc("/hello", printHello) {
			> 로그 기록 시
			> 누군가 접속 시
		> factory wrapper function 빌드 가능
			> http.HandleFunc("/hello", log(printHello)) {
			> http.HandleFunc("/bye", log(printBye)) {

	*/
}
