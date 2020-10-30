package main

import (
	"fmt"
)

func main() {
	fmt.Println("===========================================")
	declareVariables()
	fmt.Println("===========================================")
	numericTypes()
	fmt.Println("===========================================")
	stringType()
	fmt.Println("===========================================")
	runeType()
	fmt.Println("===========================================")
	arrayType()
}

func declareVariables() {
	// 선언된 변수는 반드시 사용이 되어야 하고, 사용되지 않으면 컴파일 에러 발생한다
	// var {변수명} {타입} = {값}
	var i int = 10 // literal value 10
	fmt.Println(i)

	// type inference
	var j = 20
	fmt.Println(j)

	/*
		:= operator: 함수 내에서 사용 가능
		그럼 var를 유지하는 이유?
		1. 함수 바깥에서 변수 선언 시 필요
		2. 어떤 타입인지 명시하고 싶은 경우
	*/
	k := 30
	fmt.Println(k)

	var l byte = 40
	fmt.Println(l)

	/*
		Go에서 변수는 모두 값을 가지므로, 선언만 하더라도 초기값(zero value)를 갖게 된다
	*/
	var m int
	fmt.Println(m)
	m = 50 // 이미 위에서 선언을 헀으므로 := 연산자를 사용하면 안 된다. :=은 선언(declaration)이며 할당(assignment)이 아니다.
	fmt.Println(m)
}

func numericTypes() {
	/*
		8 bits = 1 byte
		[INT]
		[signed]
		int8    -128                        127
		int16   -32,768                     32,767
		int32   -2,147,483,648              2,147,483,647
		int64   -9,223,372,036,854,775,808  9,223,372,036,854,775,807   quintillion

		[unsigned]
		uint8   0                           255
		uint16  0                           65535
		uint32  0                           4,294,967,293
		uint64  0                           18,446,744,073,709,551,615

		byte는? int8
		int는? int32 OR int64
		uint는? uint32 OR uint64


		[FLOAT]
				Number of bits              Number of Decimal Digits
		float32 32bits                      ~7
		float64 64 bits                     ~15

		!!비교 위해서는 반드시 타입 변환을 해야 한다
	*/
	var i8 int8 = 20
	var f32 float32 = 5.6
	/* fmt.Println(i8 + f32) */    // Invalid operation: i8 + f32 (mismatched types int8 and float32)
	fmt.Println(float32(i8) + f32) // 25.6
	// int8(f32)는 반올림(round)이 아니라 소수점 이하를 잘라낸다(truncated).
	fmt.Println(i8 + int8(f32)) // 25
	// 타입 변환 시 한 변수를 받지만, 표현식(expression)을 받을 수도 있다
	fmt.Println(i8 + int8(f32+1.93)) // 27
	/*
		1. 타입 간의 자동 형변환 금지
			var i32 int32= 20
			fmt.Println(i8 + i32)   // invalid operation: i8 + i32 (mismatched types int8 and int)
			var i int = 50
			fmt.Println(i + ui8)  // invalid operation: i + ui8 (mismatched types int and uint8)

		2. 사이즈가 다른(int8와 int32, 또는 int8와 uint8) 간의 자동 형변환 금지
			var ui8 uint8 = 20
			fmt.Println(i8 + ui8)  // invalid operation: i8 + ui8 (mismatched types int8 and uint8)
	*/
	var i32 int32 = 20
	fmt.Println(int32(i8) + i32)
	var ui8 uint8 = 20
	fmt.Println(uint8(i8) + ui8)

	// PHP 같은 경우 0을 false로도 보지만, Go에서는 반드시 형변환 해야 한다
	/*
		Non-bool '(0)' (type untyped int) used as condition
		if(0) {

		}
	*/

	var i_1st int = 2100000000  // 2,100,000,000 = 21억
	var i_2nd int = 42000000000 // 42,000,000,000 = 420억
	// fmt.Println(getType(i_1st)) // int
	// fmt.Println(getType(i_2nd))
	fmt.Println(i_1st + i_2nd) // 44,100,000,000
}

func stringType() {
	var s1 string // s1의 length: 0
	s1 = "Hello, World!"
	fmt.Println(s1)
	s2 := "Hello, World!"
	fmt.Println(s2)
	/*
		\: 이스케이프
		\n: 줄바꿈
		\t: 탭
	*/
	s3 := "Hello, \n\t\"World!\" with a backslash \\"
	fmt.Println(s3)
	/*
		`(backtick)은 PHP의 HEREDOC과 유사
	*/
	s4 := `
Hello, 
	"World!" with a backslash \
`
	fmt.Println(s4)
	s5 := "👋 🌍"
	fmt.Println(s5)
	s6 := s1 + " " + s5
	fmt.Println(s6)
	/*
		Go는 문자열(string)을 수정할 수 없는(unmodifiable) 일련의 바이트(sequence of bytes)로 다룬다
		문자열이 일련의 바이트이므로 Go는 문자열로부터 단일 바이트를 가져오기 쉽다
	*/
	b1 := s1[0]
	b1String := string(b1) // 72 bytes => H
	b2 := s1[4]
	b2String := string(b2)                      // 111 bytes => o
	fmt.Println(s1, b1, b1String, b2, b2String) // Hello, World! 72 H 111 o
	fmt.Println(s1, len(s1))                    // Hello, World! 13
	s1Sub1 := s1[0:5]
	fmt.Println(s1Sub1) // Hello
	s1Sub2 := s1[7:12]
	fmt.Println(s1Sub2) // World
	s1Sub3 := s1[:5]
	fmt.Println(s1Sub3) // Hello
	s1Sub4 := s1[7:]
	fmt.Println(s1Sub4) // World!
	// 만약 ASCII 코드가 아닌 문자를 다룰 경우 부분 문자열(substring)과 길이(length)를 얻는 건 좀 더 복잡해진다

	s7 := `
$s1
{$s1}
s1 %s
%d %s
`
	fmt.Println(s7)
	/*
		$s1
		{$s1}
		s1 %s
		%d %s
	*/
	fmt.Println(fmt.Sprintf(s7, "backtick 내에서 변수 사용 시작", 14, "<<<< 14가 들어간다"))
	/*
		$s1
		{$s1}
		s1 backtick 내에서 변수 사용 시작
		14 <<<< 14가 들어간다
	*/
}

func runeType() {
	/*
		rune = int32
		정수 값으로부터 문자 값을 구별하기 위해 사용
		rune > string 타입 변환 가능
		Go는 문자열의 모든 rune을 순서에 따라 순회하는 방식을 내장
	*/
	var r1 rune = 127757
	s1 := "Hello, " + string(r1)
	fmt.Println(s1)
	r2 := '🌍' // rune = int32
	s2 := "Hello, " + string(r2)
	fmt.Println(s2)
}

func arrayType() {
	var vals1 [3]int
	vals1[0] = 2
	vals1[1] = 4
	vals1[2] = 6
	fmt.Println(vals1, vals1[0], vals1[1], vals1[2]) // [2 4 6] 2 4 6
	/*
		배열의 길이는 array 타입의 일부로 고려된다
		[3]int를 [4]int에 할당할 수 없으며, 숫자 타입과 다르게 conversion 방법이 없다.
		따라서 크기를 동적으로 증가시키거나 배열을 자르는 등의 기능을 가지고 있지 않다

		var vals2 [4]int = vals1
		fmt.Println(vals1, vals2)  // cannot use vals1 (type [3]int) as type [4]int in assignment
	*/

	var vals3 []int
	fmt.Println(vals3) // []
	// fmt.Println(vals3, vals3[1])  //  runtime error: index out of range [1] with length 0
	/*
		vals3[1] = 1
		fmt.Println(vals3)  // panic: runtime error: index out of range [1] with length 0
	*/
}
