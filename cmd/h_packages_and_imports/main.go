package main

import (
	"fmt"
	"math"
	"math/rand" // math 안에 rand가 있다고 하더라도, 명시적으로 import 해야 한다
	"strings"
	"time"
	"unicode/utf8"
)

/*
	Package?
	> 한 영역의 기능(functionality)에 집중
	> 대부분의 Go 프로그램은 여러 패지키로 쪼개질 수 있다
	> Go 프로그램에서 각 디렉토리는 유니크한 패키지와 연관(associated)된다
	> 크게 내가 생성한 패키지, 남이 생성한 패키지, 표준 패키지 라이브러리 세 종류의 패키지 존재

	Import?
	> 패키지가 import 되지 않으면 패키지의 코드에 접근할 수 없다
	> 이는 자바와 다른데, 패키지를 import 한다는 것은 단순히 더 짧은 이름으로 그 안의 코드를 참조할 수 있다는 것을 의미
	> Java와 Python에서는 패키지에서 한 항목만 임포트할 수 있지만, Go에서는 이를 허용하지 않으며, 패키지 임포트 하면 패키지의 모든 것에 대한 접근 가능
*/

/*
	Go는 내장 타입에 메서드를 두지 않고, 해당 기능을 제공하기 위해 표준 라이브러리에 패키지를 가지고 있다
	string package에는 UTF-8로 싱행되는 함수 없으며, 해당 함수는 Unicode/UTF-8 패키지에 있다
	가령 len의 경우 문자열의 바이트의 수를 반환하며, 실제 문자열의 문자 수를 알고 싶다면 UTF-8 RuneCountInString 함수 사용
*/
func stringsPkgs() {
	s1 := "This is a test 123 😁"
	/*
		Map 메서드는 mapping 함수에 따라 수정된 문자들의 복사본을 반환단다
		만약 mapping이 negative value를 반환하면, 해당 문자가 치환 없이 문자열에서 누락됨을 의미
	*/
	s2 := strings.Map(rot13, s1)
	fmt.Println(s2) // Guvf vf n grfg 123 �
	s3 := strings.Map(rot13, s2)
	fmt.Println(s3) // This is a test 123 �

}

// rot13(Rotate by 13)은 암호(encryption) 함수로, 영어 알파벳의 문자를 13으로 회전(rotate, 즉 밀어서) 시켜서 만든다
// https://ko.wikipedia.org/wiki/ROT13
func rot13(input rune) rune {
	// '문자' 처럼 홑따옴표 안에 문자를 위치시키면 해당 문자의 rune 값을 얻을 수 있다
	if input >= 'A' && input <= 'Z' {
		/*
			[대문자인 경우]
			(input -'A'): 'A'에서 시작하므로 0부터 최대 25
			(0 ~ 25) + 13: 13만큼 밀어낸다
			((0 ~ 25) + 13) % 26: 알파벳의 개수(26)으로 나눈 나머지로 0부터 25 사이의 값 안에서 회전(rotate)하게 된다
			(((0 ~ 25) + 13) % 26) + 'A'(65): 위에서 연산된 값을 'A'(대문자 알파벳의 시작값)에 더하면 회전 된 알파벳의 rune 반환한다
		*/
		return (((input - 'A') + 13) % 26) + 'A'
	}
	if input >= 'a' && input <= 'z' {
		return (((input - 'a') + 13) % 26) + 'a'
	}

	return input
}

func unicodeAndUtf8Pkgs() {
	s1 := "😀 😑 🍟"
	fmt.Println(s1)
	fmt.Println(len(s1))                    // 14: number of bytes of string
	fmt.Println(utf8.RuneCountInString(s1)) // 5: number of characters in string
}

func mathPkgs() {
	/*
		func Cos(x float64) float64
		func Exp(x float64) float64
		func Log(x float64) float64
		func Min(x float64) float64
		func Max(x float64) float64
		모두 float64로 계산된다
		정수 연산하고 싶다면,
		1. 파라미터로 넘길 때 float64()로 형변환
		2. 연산
		3. 결과값을 int()로 형변환
	*/
	/*
		pseudo-random number generator
		func Seed(seed int64)  // 같은 값을 넘기면, 매번 같은 순서의 숫자들을 반환받게 된다. 무작위 수 원하면 time의 unix nano 값을 넘긴다
		func Intn(n int) int  // 0보다 크고 전달되는 int n보다 작은 무작위 수 반환
	*/
	// 호출 시마다 무작위 수를 반환하기 위해 씨앗(seed)으로 int64 나노초 unix time을 사용
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(10)
	r2 := rand.Intn(10)
	a := int(math.Max(float64(r1), float64(r2)))
	fmt.Println(r1, r2, a)
}

func main() {
	fmt.Println("===========stringsPkgs===========")
	stringsPkgs()
	fmt.Println("===========unicodeAndUtf8Pkgs===========")
	unicodeAndUtf8Pkgs()
	fmt.Println("===========mathPkgs===========")
	mathPkgs()
}
