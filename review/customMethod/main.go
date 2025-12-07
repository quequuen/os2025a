package main

import "fmt"

// 메서드를 main 안에 정의할 수 없어서 꼭 밖에 정의해야 함
type Grade float64

// 임베딩

type Student struct{
	Name  string
	Grade
}
	
func (g Grade) isPassing() bool {
	return g >= 60.0
}


func main(){


	// 사용자 정의 메서드 예제
	// var myGrade Grade = 75.5
	myGrade := Grade(75.5)
	if myGrade.isPassing() {
		fmt.Println("Passing")
	} else {
		fmt.Println("Failing")
	}
	
}
