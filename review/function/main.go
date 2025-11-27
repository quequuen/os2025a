package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

//다중 반환 예제2
func floatParts(number float64)(integerPart int, fractionalPart float64){
	// 이런 식으로 문서화를 위해 반환 값에 이름을 부여하는 것도 가능
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

//다중 반환 예제
func manyReturns()(int, bool, string) {
	return 1, true, "hello"
}

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
func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is incalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	// fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / 10.0, nil //페인트의 양과 함께 아무 에러도 없음을 나타내는 "nil" 값을 반환.

}

func main(){


	//다중 반환 예제2
	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)

	// 다중 반환 예제
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)
	

	// 에러 생성과 반환
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())
	// log.Fatal()
	//에러를 출력하고 빠져나오는 메소드. 밑에 코드는 작동하지 못함

	// 페인트 양 누적 및 계산 함수
	var amount, total float64
	amount, err = paintNeeded(4.2, 3.0)
	// 항상 에러를 처리할 것!!
	// fmt.Println(err)
	// fmt.Printf("%0.2f liters needed\n", amount)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%0.2f liters needed\n", amount)
	}
	total += amount
	amount, err = paintNeeded(5.2, 3.5)
	// fmt.Println(err)
	// fmt.Printf("%0.2f liters needed\n", amount)
	if err != nil {
		log.Fatal()
		//보통 이런 식으로 nil이 아니라면 에러를 찍고 프로그램을 강제 종료함
	}else{
		fmt.Printf("%0.2f liters needed\n", amount)
	}
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
