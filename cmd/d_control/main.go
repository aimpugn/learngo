package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("==================ifElse======================")
	ifElse()
	fmt.Println("==================forLoop======================")
	forLoop()
	fmt.Println("==================switchStmt======================")
	// switchStmt()
}

func ifElse() {
	a := 10
	if a > 5 { // 괄호로 감싸지 않아도 된다
		// Go는 block scope, 따라서 b를 if문 밖에서 사용할 수 없다
		b := a / 2
		fmt.Println("a > 5: ", a, b)
	} else {
		/*
			}
			else { <<- 이렇게 작성은 불가능
		*/
		fmt.Println("a < 5: ")
	}
	// fmt.Println("a > 5: ", a , b)  // undefined: b

	/*
		if (0) {  // non-bool 0 (type int) used as if condition
			fmt.Println("when 0")
		}
	*/

	/*
		if a = 5 {  // syntax error: assignment a = 5 used as value
		}
	*/

	// 비교 전에 새로운 변수 선언하여 if/else 블록 내에서 사용 가능
	if b := a / 2; b > 5 {
		fmt.Println("b > 5: ", a, b) // b > 5:  10 5
	} else {
		fmt.Println("b > 5: ", a, b)
	}
	// fmt.Println("a > 5: ", a , b)  // undefined: b
}

func forLoop() {
	fmt.Println("=================1st=======================")
	for i := 0; i < 10; i++ {
		// 초기값을 한번 선언하고 그 다음에 "="로 변수에 값을 할당
		// var delimiter string
		delimiter := ""
		if i == 9 {
			delimiter = "\n"
		} else {
			delimiter = ", "
		}
		fmt.Print(strconv.Itoa(i) + delimiter)
	}

	fmt.Println("=================2nd:break=======================")
	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("=================3rd:continue=======================")
	a := 3
	for i := 0; i < 10; i++ {
		if i == a {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("=================4th:while=======================")
	// while이 따로 없고, for를 사용하면 된다
	j := 0
	for j < 10 {
		fmt.Println(j)
		j += 1
	}
	fmt.Println(j)

	fmt.Println("=================4th:while true=======================")
	j = 0
	for {
		j++
		if j > 100 {
			break
		}
	}
	fmt.Println(j) // 101

	fmt.Println("=================5th:for range=======================")
	s1 := "Hello, World!"
	// multibyte 문자를 사용할 경우, 유효한 rune을 얻는 한 방법
	for offset, valueInRune := range s1 {
		fmt.Println(offset, valueInRune, string(valueInRune))
		/*
			0 72 H
			1 101 e
			2 108 l
			3 108 l
			4 111 o
			5 44 ,
			6 32
			7 87 W
			8 111 o
			9 114 r
			10 108 l
			11 100 d
			12 33 !
		*/
	}
	s2 := "👋 🌍"
	// multibyte 문자를 사용할 경우, 유효한 rune을 얻는 한 방법
	for offset, valueInRune := range s2 {
		fmt.Println(offset, valueInRune, string(valueInRune))
		/*
			offset은 문자가 아닌 byte 단위로 계산이 된다
			첫 👋의 문자가 4 bytes, 빈 공백이 1 byte, 그 다음 🌍이 4bytes
			0 128075 �
			4 32
			5 127757 �
		*/
	}
}

func switchStmt() {
	// 하나 또는 그 이상 else if를 사용한다면 instead 고려
	// 커맨드 라인으로 전달되는 인자(argument) 사용
	word := os.Args[1]
	/*
		if word == "hello" {
			fmt.Println("Hi yourself")  // go run main.go hello
		} else if word == "goodbye" {
			fmt.Println("So long!")  // go run main.go goodbye
		} else if word == "greetings" {
			fmt.Println("Salutations!")  // go run main.go greetings
		} else {
			fmt.Println("I don't know what you said")  // go run main.go what?
		}
	*/

	/*
		Go에서는 단일 케이스에 대해서만 코드가 실행되며, break가 필요 없다
		단, switch 문에서 비교하려는 값의 타입과 case 문장의 비교되는 값의 타입이 같아야 한다
	*/
	fmt.Println("=================1th:switch=======================")
	greet := "greetings"
	// if 문에서처럼, switch 내에서만 유효한 변수 선언 가능
	switch l := len(word); word { //
	case "hi":
		fmt.Println("Very informal!")
		// C나 Java에서처럼 단일 케이스에서 끝나지 않고 다음 케이스로 넘어가게 하고 싶을 경우 사용
		fallthrough
		/*
			Very informal!
			Hi yourself
		*/
	case "hello":
		fmt.Println("Hi yourself") // go run main.go hello
	case "goodbye", "bye": // 여러 값에 대해 같은 결과를 출력하고 싶으면, 콤마로 구별하여 여러 값을 나열한다
		fmt.Println("So long!") // go run main.go goodbye
	case "farewell":
		// 아무 코드도 실행되지 않는다
	case greet:
		/*
			C나 Java와 달리 Go에서는 비교하려는 값이 굳이 constant value일 필요 없으며,
			어떤 타입의 값이든 비교 가능하다. 변수(variable), 상수(constant) 또는 literal value 모두 비교 가능
		*/
		fmt.Println("Salutations!") // go run main.go greetings
	default:
		fmt.Println("I don't know what you said, but iw was", l, "characters length") // go run main.go what?
		// I don't know what you said, but iw was 5 characters length
	}

	fmt.Println("=================2nd:switch like if-else=======================")
	c := "crackerjack"
	switch l := len(word); { // case에서 비교를 하기 위해 word를 switch 문 조건에서는 제거. 이때 비교할 변수가 없어도 블록 스코프 변수 선언 시 세미콜론(;) 필수
	case word == "hi":
		fmt.Println("Very informal!")
		fallthrough
	case word == "hello":
		fmt.Println("Hi yourself")
	case l == 1:
		fmt.Println("I don't know any one letter words")
	case 1 < l && l < 10, word == c:
		fmt.Println("This word is either", c, "or it is 2-9 characters long")
	default:
		fmt.Println("I don't know what you said, but iw was", l, "characters length") // go run main.go what?
	}
}
