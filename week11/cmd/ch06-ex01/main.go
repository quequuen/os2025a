package main

import "fmt"

func main(){
	// 슬라이스 리터럴 선언 및 할당
	// var subjects []string 
	// subjects = make([]string,3)
	subjects := []string{"Go", "", "Python"}
	

	for _, subject := range subjects {
		fmt.Println(subject)
	}
	
}