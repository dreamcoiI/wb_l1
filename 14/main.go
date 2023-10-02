package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num interface{} = 1
	var str interface{} = "str"
	var b interface{} = true

	ch := make(chan int)

	chIn := interface{}(ch)

	fmt.Println(reflect.TypeOf(num).String())
	fmt.Println(reflect.TypeOf(str).String())
	fmt.Println(reflect.TypeOf(b).String())
	fmt.Println(reflect.TypeOf(chIn).String())
}
