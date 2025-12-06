package main

import "fmt"



func main(){

	


	//내부 배열 변경
	array1 := [5]string{"A","B","C","D","E"}
	slice1 := array1[1:4] // 배열의 일부를 참조하는 슬라이스 생성
	slice1[0] = "X"       // 슬라이스를 통해 내부 배열 변경
	fmt.Println(array1)         // [A X C D E] 출력

	//슬라이스 확장
	slice2 := array1[2:4] // 배열의 일부를 참조하는 슬라이스 생성
	slice2 = append(slice2, "Y") // 슬라이스 확장
	fmt.Println(slice2)                // [C D Y] 출력
	fmt.Println(array1)                // [A X C D Y] 출력, 내부 배열도 변경됨

	//슬라이스 중복 내부 배열 변경
	slice3 := array1[1:4] // 배열의 일부를 참조하는 슬라이스 생성
	slice4 := array1[2:5] // 배열의 일부를 참조하는 슬라이스 생성
	slice3[1] = "Z"       // 슬라이스3을 통해 내부 배열 변경
	fmt.Println(array1)         // [A X Z D Y] 출력
	fmt.Println(slice3)         // [X Z D] 출력
	fmt.Println(slice4)         // [Z D Y] 출력, 슬라이스4도 변경됨

	// 슬라이스 리터럴
	notesliteral := []string{"one","two","three","four","five","six","seven"}
	for i := range notesliteral{
		fmt.Println(notesliteral[i])
	}



	// 슬라이스
	// 슬라이스는 배열과 다르게 크기가 고정되어 있지 않다.
	// 동적으로 크기가 변할 수 있다.
	// 슬라이스는 배열의 일부를 '참조'하는 자료구조이다.

	// var notes []string // 배열과 달리 크기를 지정하지 않음!
	// notes = make([]string, 7)
	//일곱 개의 문자열을 담을 수 있는 슬라이스 생성
	notes := make([]string, 7)
	notes[0] = "zero"
	notes[1] = "one"
	notes[2] = "two"
	print(notes[0])
	print(notes[1])
	print(notes[2])
	print(notes)
	
} 