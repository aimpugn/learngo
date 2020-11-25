package main

import (
	"fmt"
)

/*
	C, C++에서는 메모리 누수와 프로그램이 깨지지 않도록 포인터 사용 시 세심하게 관리해야 했지만,
	Go에서는 C, C++에서처럼 포인터를 사용할 수 없기에 두려워 할 필요가 없다!

    Go에서는 자체적인 내장(built-in) 배열과 문자열 타입이 있으므로 이를 시뮬레이션(simulate)하는 데 포인터를 사용하지 않는다
		> 버퍼 오버플로우 에러를 없는 데 도움을 준다
	Go는 프로그램에서 생성된 모든 포인터를 추적하며, 가비지 컬렉션을 사용하여 할당 해제(deallocate) 하여 더이상 사용되지 않게 한다
		> 메모리 누수를 방지하고 우연한(accidental) 할당 해제(deallocation)를 방지하여 random crash 방지

	또한 Go는 강한 타입(strongly typed)의 언어
		> 메모리를 할당하여 원하는 타입으로 사용할 수 없다
		> 어떤 포인터 타입을 다른 타입으로 캐스팅할 수 없다
*/

func reference() {
	a := 10
	/*
		&(ampersand) 사용하면 참조(reference) 의미
		여기서는 a가 저장된 메모리의 위치(hexadecimal)를 나타낸다

		*(asterisk) 사용하면 역참조(dereference) 하여 해당 메모리에 저장된 값을 볼 수 있다
	*/
	b := &a
	c := a // copy of the value of a
	// a  hexadecimal
	fmt.Println(a, b, *b, c) // 10 0xc000014078 10 10

	a = 20
	fmt.Println(a, b, *b, c) // 20 0xc000014078 20 10

	*b = 30
	fmt.Println(a, b, *b, c) // 30 0xc000014078 30 10

	c = 40
	fmt.Println(a, b, *b, c) // 30 0xc000014078 30 40
}

func typePointer() {
	/*
				// 이 포인터에 초기값(zero value) 할당된다
				// 어떤 타입을 가리키든, 모든 포인터는 같은 초기값(=nil)을 갖으며, 읽거나 쓸 수 없다
				// nil은 메모리 위치가 없으며, 1의 부재(absence of 1)
				var b *int
		        fmt.Println(b)
				fmt.Println(*b)  // panic: runtime error: invalid memory address or nil pointer dereference
				// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48d378]
	*/
	// 포인터를 만들기 위해 Go에 내장된 함수 사용 가능.
	b := new(int) // 포인터 생성 + 메모리 할당

	fmt.Println(b)  // 0xc0000a6078 -> hexadecimal 메모리 위치
	fmt.Println(*b) // 초기값 0
}

//              int에 대한 포인터로, hexadecimal 값이 들어와야 한다
func setTo10ByPointer(a *int) {
	// 역참조: int에 대한 포인터가 가리키는 값에 10 할당
	*a = 10
}

func interveneSetTo10ByPointer(a *int) {
	a = new(int) // 새로운 포인터 생성
	*a = 10      // 새로운 포인터의 메모리에 10 저장
}

func reference1(a *[]int, b *int) {
	tmp := *a
	fmt.Println("a:", a, "*a:", *a, "tmp:", tmp, "b:", b, "*b:", *b)
	reference2(a, b)
}

func reference2(c *[]int, d *int) {
	tmp := *c
	tmp[0], tmp[1] = tmp[1], tmp[0]
	fmt.Println("a:", c, "*c:", *c, "tmp:", tmp, "d:", d, "*d:", *d)
}

func main() {
	fmt.Println("=================reference=================")
	reference()
	fmt.Println("=================typePointer=================")
	typePointer()
	fmt.Println("=================setTo10ByPointer=================")
	a := 20
	fmt.Println(a) // 20
	//          a에 대한 포인터(hexadecimal 메모리 위치)를 전달
	setTo10ByPointer(&a)
	fmt.Println(a) // 10
	fmt.Println("=================interveneSetTo10ByPointer=================")
	b := 20
	fmt.Println(b) // 20
	//          a에 대한 포인터(hexadecimal 메모리 위치)를 전달
	interveneSetTo10ByPointer(&b)
	fmt.Println(b) // 10
	fmt.Println("=================nested reference=================")
	arr := []int{
		1,
		2,
		3,
		4,
		5,
	}
	a1 := 2
	reference1(&arr, &a1)
}
