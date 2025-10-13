package main

import (
	"fmt"
	"reflect"
)

func main() {
	//명시적 casting
	//Go는 엄격한 타입 언어이기 때문에 다른 타입 간의 연산이 불가능
	//따라서, 다른 타입 간의 연산을 위해서는 명시적 형 변환이 필요
	//형 변환은 함수 형태로 제공됨
	//형 변환 함수: int(), float64(), string() 등
	
	var length float64 = 1.2
	var width int = 2
	// fmt.Println("Error Example", length*width) // 타입 불일치로 에러 발생
	fmt.Println("Area is", length*float64(width)) // 일시적 형변환으로 인한 에러 해결
	//원본 자체는 바뀌지 않음
	fmt.Println("면적은", int(length)*width)
	fmt.Println("형 변환", reflect.TypeOf(float64(width)))
	fmt.Println("length > width?", length > float64(width))
}
