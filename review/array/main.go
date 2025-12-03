package main

import "fmt"

// 배열
// 배열의 번호: 인덱스(index)
// 배열은 선언한 크기가 고정되어 변하지 않는다.
func main(){
	


	// 배열의 범위를 벗어난 인덱스를 조회하려고 하면 패닉 발생

	// 배열 리터럴
	// 선언과 동시에 초기값을 다른 값으로 할당
	var notesliteral [7]string = [7]string{"do","re","mi"}
	fmt.Println(notesliteral)
	for _, note := range notesliteral{
		fmt.Println(note)
	}


	// 기본 문자열 배열
	// 선언과 동시에 초기값들로 채워짐. -> 추후에 하나씩 채움
	var notes [7]string
	notes[0] = "do"

	fmt.Println(notes) //[do      ]
}