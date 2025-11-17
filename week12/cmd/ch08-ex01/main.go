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
	// s1.name = "Kim Inha"
	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
}
