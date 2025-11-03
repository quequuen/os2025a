package main

import "fmt"

func main(){
	//배열(선언 후 할당)
	var arrayBool [3]bool
	var arrayInt [3]int

	fmt.Println(arrayBool[1])	//zero value
	arrayInt[1]++
	arrayInt[1]++
	fmt.Println(arrayInt[1])	//zero value + 2

	//array 리터럴(선언과 동시에 할당)
	var arrayBool2 [3]bool = [3]bool{true, false, false}
	fmt.Println(arrayBool2)
	fmt.Printf("%#v\n", arrayBool2)

	arrayInt2 := [3]int{1,2,3}
	fmt.Println(arrayInt2)
	fmt.Printf("%#v\n", arrayInt2)

	//반복문
	arrayBool3 := [3]bool{true, false, true} 
	arrayInt3 := [3]int{-9, 11, 7}

	for i := 0; i < len(arrayInt3); i++ {
		fmt.Println(i, arrayInt3[i])
		fmt.Println(i, arrayBool3[i])
	}

	//반복문 range 쓰는 예제 
	// numbers := [3]int{-9, 11, 7}

	// for i := 0; i < len(arrayInt3); i++ {
	// 	fmt.Println(i, numbers[i])
	// }

	numbers := [3]int{-9, 11, 7}

	for i, number := range numbers {
		fmt.Println(i, number)
	}


}
