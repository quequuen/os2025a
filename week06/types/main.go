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
		// 단축 변수 선언문(short variable declartion)
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
	{
		//Naming rules
		
		// 1. 숫자로 시작할 수 없다. (영어로 시작해야만 함.)
		// var 64f float64

		// 2. 추가적인 문자나 숫자의 제약이 없다.
		
		// 3. 대소문자 구분을 분명히 한다.
		// fmt.println()
		// +) 앞 글자를 소문자로 하게 되면 외부에서 접근이 불가능하다. 강제 static화..?

		// 4. 관례상 변수명이나 함수명은 카멜케이스로 해라

		// 5. 최대한 생략해라
		//  index -> i
	}
	
}