package main

import "fmt"

func main(){
	// 슬라이스 리터럴 선언 및 할당
	// var subjects []string 
	// subjects = make([]string,3)
	subjects := []string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[1:3]

	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("----슬라이스 일부만 출력----")
	for i:=0; i<len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
	
}