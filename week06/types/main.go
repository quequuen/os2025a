package main

import (
	"fmt"
	"reflect"
	// "math"
	// "strings"
)

func main(){
	{
		name:="Kim Inha"
		// ':='는 선언할 때 쓰는 거기 때문에 재할당 용도로는 쓰지 못함.
		// 이것도 할당임

		// var name string
		// name = "Kim Inha"
		// 명확하게 타입을 알려주는 방식

		// var name = "Kim Inha"
		// 타입 추론 방식
		fmt.Println(name, reflect.TypeOf(name))


		// fmt.Println(math.Floor(3.14));
		// fmt.Println(strings.Title("head first go"));
		// fmt.Println(math.Round(2.31));
	}
	{
		//zero value 
		var f64 float64
		var str string
		var i32 int
		var b bool

		fmt.Println(f64, reflect.TypeOf(f64))
		fmt.Println(str, reflect.TypeOf(str))
		fmt.Println(i32, reflect.TypeOf(i32))
		fmt.Println(b, reflect.TypeOf(b))
	}
	
}