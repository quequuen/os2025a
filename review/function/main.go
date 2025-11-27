package main

import (
	"errors"
	"fmt"
	"log"
)

//return error 예제 함수
// func double(number float64) float64{
// 	return number *2
// 	fmt.Println(number*2)
// }
// return 문 밑에 다른 코드가 있으면 에러가 뜸!

//두배 계산 함수
// func double(number float64) float64{
// 	// 항상 변수 타입 이 규칙으로 써야 함
// 	// 리턴값도 뒤에 써줘야 함
// 	return number*2
// }

//페인트 양 계산 함수
func paintNeeded(width float64, height float64) float64{
	area := width * height
	// fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / 10.0

}

func main(){

	
	

	// 에러 생성과 반환
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())
	log.Fatal()
	//에러를 출력하고 빠져나오는 메소드. 밑에 코드는 작동하지 못함

	// 페인트 양 누적 및 계산 함수
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)


	// 두배 계산 함수
	// dozen := double(6.0)
	// fmt.Println(dozen)
	// fmt.Println(double(4.2))


	//페인트 양 계산 함수
	// paintNeeded(4.2, 3.0)
	// paintNeeded(5.2, 3.5)
	// paintNeeded(5.0, 3.3)

	
}
