package main

import "fmt"

func main(){
	// 슬라이스 리터럴 선언 및 할당
	// var subjects []string 
	// subjects = make([]string,3)
	subjects := []string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[1:3]
	// 슬라이스는 범위를 넘을 수 없다!
	// subjectsSlice := subjects[1:4] //에러 발생

	fmt.Println("----슬라이스 전체 출력----")
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("----슬라이스 일부만 출력----")
	for i:=0; i<len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}

	// 배열 리터럴을 슬라이싱 하면 슬라이스 리터럴이 됨.(배열을 이용해 슬라이스 생성 가능)
	// 해당 동작은 깊은 복사가 아닌 얕은 복사로 값이 바뀌면 그 값을 공유함.(포인터 개념)
	array3:= [5]string{"C", "C++", "Java", "HTML", "CSS"}
	arraySlice := array3[2:4]
	array3[2] = "ChangedJava"
	arraySlice[0] = "ChangedAgainJAVA"
	
	fmt.Println(array3)
	fmt.Println(arraySlice)

	// go의 append는 메소드가 아닌 내장 함수로 동작함.
	arraySlice = append(arraySlice, "GoLang", "Python", "Cotlin", "Git", "Database")
	for _, v := range arraySlice{
		fmt.Println(v)
	}
	
}