package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("===========================================")
	ifElse()
	fmt.Println("===========================================")
	forLoop()
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

	fmt.Println("=================2nd=======================")
	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("=================3rd=======================")
	a := 3
	for i := 0; i < 10; i++ {
		if i == a {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("=================4th=======================")
	// while이 따로 없고, for를 사용하면 된다
	j := 0
	for j < 10 {
		fmt.Println(j)
		j += 1
	}
	fmt.Println(j)

	j = 0
	for {
		j++
		if j > 100 {
			break
		}
	}
	fmt.Println(j) // 101

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
