package i_user_defined_pkgs

import (
	"strings"
	"unicode/utf8"
)

/*
	https://www.digitalocean.com/community/tutorials/understanding-package-visibility-in-go
	패키지에서 외부에서 사용할 수 있는 함수는 대문자로 시작한다
	아래의 경우 default_char 변수는 소문자로 시작하여 패키지에서만 사용할 수 있으며
	Format, FormatRune 함수는 대문자로 시작하여 패키지 외부에서도 사용할 수 있는 public 함수다
*/
// 패키지 레벨의 변수
var default_char = ' '

/*
	Format 함수는 문자열과 정수를 매개변수로 받으며, 전달된 정수 크기만큼 좌측 공백이 패딩된 문자열을 반환
	문자열의 길이가 지정된 길이보다 더 큰 경우, 원래 문자열이 반환된다
*/
func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

/*
	FormatRune 함수는 문자열, 정수, 룬을 매개변수로 받고 정수 길이만큼 지정된 rune을 좌측 패딩한 문자열을 반환한다
	만약 문자열의 길이가 지정된 길이보다 큰 경우, 원래 문자열이 반환된다
*/
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
