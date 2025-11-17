package main

import "fmt"

type subscriber struct{
	name string
	price int
}

func applyPrice(s* subscriber){
	s.name = "Kim Init"
	s.price= 10000
}

func main(){
	var s1 subscriber
	var p *subscriber = &s1
	// 주소를 줘야 포인터를 이용해서 호출 가능
	// s1.name = "Kim Inha"
	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
	// 밑에 두 방법 다 가능
	fmt.Println((*p).price)
	fmt.Println(p.price)
}
