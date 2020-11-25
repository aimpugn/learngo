package main

import (
	fmt "fmt"
	"strconv"
)

/*
rune
*/
func main() {
	fmt.Println("'{'", '{')
	fmt.Println("'}'", '}')
	fmt.Println("'['", '[')
	fmt.Println("']'", ']')
	/*
		func isSpace(c byte) bool {
			return c <= ' ' && (c == ' ' || c == '\t' || c == '\r' || c == '\n')
		}
	*/
	fmt.Println("' '", ' ')
	fmt.Println("'\\t'", '\t')
	fmt.Println("'\\r'", '\r')
	fmt.Println("'\\n'", '\n')
	fmt.Println("string(30)", string(30))
	fmt.Println("string(31)", string(31))
	fmt.Println("string(32)", string(32))
	fmt.Println("strconv.FormatInt(32, 16)", strconv.FormatInt(32, 16))
	fmt.Println("string(33)", string(33))
	fmt.Println("string(34)", string(34))
	fmt.Println("string(35)", string(35))
	fmt.Println("string(0x20)", string(0x20)) // 32의 hexdecimal 값
	fmt.Println("string(0x21)", string(0x21))
	fmt.Println("Hello, World")
	fmt.Println([]rune("Hello, World"))
}
