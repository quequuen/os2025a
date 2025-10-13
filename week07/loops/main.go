package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	{
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

	{
	//time 패키지
	var now time.Time = time.Now()
	var year int = now.Year()
	var month time.Month = now.Month()
	//Month는 int형이 아니기 때문에 time.Month()로 변환해야 함
	fmt.Println(year)
	fmt.Println(month)

	//예제 1
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	//예제 2
	univ:= "Go Inha$"
	changer:= strings.NewReplacer("$", "!")
	changed:= changer.Replace(univ)
	fmt.Println(changed)
	fmt.Println(univ) //원본은 바뀌지 않음
	}
}