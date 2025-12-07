package main

import "fmt"

//포인터 사용한 구조체 예제
type subscriber struct{
	name  string
	price int
}

func applyPrice(s *subscriber){
	s.price = 10000
	s.name = "Bob"
}

type Person struct {
	Name string
	Age  int
	City string
}

func main(){
	// 구조체 예제
	var s1 subscriber
	var p *subscriber = &s1
	applyPrice(&s1)

	fmt.Println(s1.name, s1.price)

	fmt.Println((*p).name, (*p).price)
	// *p.price는 에러 발생, 우선순위 때문에 (*p).price로 작성해야 함
	fmt.Println(p.name, p.price) //포인터 변수도 .연산자로 접근 가능

	// 구조체(structure)
	person := Person{
		Name: "Alice",
		Age:  30,
		City: "Seoul",
	}

	println("Name:", person.Name)
	println("Age:", person.Age)
	println("City:", person.City)
}