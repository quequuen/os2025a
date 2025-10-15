package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// var now time.Time = time.Now()
	var year int = time.Now().Year()
	fmt.Println(year)

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// os.Stdin -> 표준 입력(키보드)
	// bufio.NewReader() -> 버퍼로 감싸서 한 줄 전체를 편하게 읽을 수 있게 해줌
	// reader -> 사용자 입력을 한 줄씩 읽는 기능이 있는 객체
	input, _ := reader.ReadString('\n') 
	// ReadString() -> 지정한 구분자가 나올 때까지 입력을 읽음
	// ReadString은 (string, error) 두 가지를 반환하지만,
	// _를 써서 에러는 무시하고 문자열만 받음-> 보통은 에러 처리를 따로 함.
	input1, err := reader.ReadString('\n')
	//error가 아닌 err라고 쓰는 이유: 이름 섀도잉을 피하기 위해서
	//error는 타입 이름이기 때문에 변수 이름으로 쓰면 안됨
	//따라서 err라고 씀
	if err != nil {
		log.Fatal(err)
		//Fatal은 err가 있든 없든 프로그램을 종료시킴
		//그래서 쓰려면 if문 안에서 써야 함
  	}
	// err가 nil이 아니면 에러가 발생한 것
	fmt.Println(input)
	fmt.Println(input1)
	fmt.Println(err)
	// nil이 뜨면 에러가 없다는 뜻

	input = strings.TrimSpace(input)
	// TrimSpace() -> 문자열의 앞뒤 공백 문자 제거, 개행(\n) 문자 제거
	grade, err := strconv.ParseFloat(input , 64)

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status := "passing"
		fmt.Println(grade)
		fmt.Println("Status is", status)
	} else{
		status := "failing"
		fmt.Println(grade)
		fmt.Println("Status is", status)
	}
	


}