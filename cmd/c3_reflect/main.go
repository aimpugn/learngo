package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := Test{}
	t.slice = t.slice[0:0]
	fmt.Println(t)
	fmt.Println(len(t.slice))
	// n := len(t.slice)
	fmt.Println("reflect.ValueOf(t)", reflect.ValueOf(t))
	fmt.Println("reflect.ValueOf(t).Kind()", reflect.ValueOf(t).Kind())
	fmt.Println("reflect.Ptr", reflect.Ptr)
	refelction(t)
	refelctionPtr(&t)
}

type Test struct {
	slice []int
}

func refelction(i interface{}) {
	rv := reflect.ValueOf(i)
	fmt.Println("refelction > rv", rv)
	fmt.Println("refelction > rv.Kind()", rv.Kind())
	fmt.Printf("refelction > rv.Kind() != reflect.Ptr is %t\n", rv.Kind() != reflect.Ptr)
}

func refelctionPtr(i interface{}) {
	rv := reflect.ValueOf(i)
	fmt.Println("refelctionPtr > rv", rv)
	fmt.Println("refelctionPtr > rv.Kind()", rv.Kind())
	fmt.Printf("refelctionPtr > rv.Kind() != reflect.Ptr is %t\n", rv.Kind() != reflect.Ptr)
}
