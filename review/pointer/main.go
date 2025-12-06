package main

import (
	"fmt"
	"reflect"
)

func swap(first *int, second *int){
	// int가 저장된 주소를 매개변수로 받음
	temp := *first
	*first = *second
	*second = temp

	fmt.Println(*first, *second)
}

func main(){

	// & = “주소 줘”
	// * = “값 줘”

	a,b := 10, 20
	fmt.Println(a,b)

	swap(&a, &b)
	fmt.Println(a, b)

	// 포인터 기본 사용법
	{
		amount := 6
		fmt.Println(amount)
		fmt.Println(&amount)
	
		var myInt int
		fmt.Println(&myInt)
		var myFloat float64
		fmt.Println(&myFloat)
		var myBool bool
		fmt.Println(&myBool)
	}

	//reflect로 타입 확인
	{
		var myInt int
		fmt.Println(reflect.TypeOf(&myInt))
		var myFloat float64
		fmt.Println(reflect.TypeOf(&myFloat))
		var myBool bool
		fmt.Println(reflect.TypeOf(&myBool))
		// *int
		// *float64
		// *bool
		// 각각의 포인터를 가져와 포인터 타입을 출력
	}

	// 포인터 변수 선언
	// 한 가지 타입에 대한 포인터 값만 가질 수 있음

	var myInt int
	// var myIntPointer *int = &myInt
	myIntPointer := &myInt
	fmt.Println(myIntPointer)

	var myFloat float64
	myFloatPointer := &myFloat
	// 포인터 변수도 마찬가지로 단축 변수 선언 사용 가능
	fmt.Println(*myFloatPointer)
	// & -> 주소를 줄 때 사용
	// * -> 해당 주소의 값을 가져올 때 사용


}
