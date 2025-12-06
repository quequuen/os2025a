package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	fmt.Printf("%4.2f \n", 1231.123123)

	data, err := os.Open("./file/data.txt")
	// 파일 열기(가져오기)
	// cmd 실행경로를 기준으로 상대경로 작성(당시 경로: review)

	var numbers [3]float64
	if err != nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)
	// 파일을 한 줄씩 읽기 위한 스캐너 생성
	
	i := 0 
	for scanner.Scan(){
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		// fmt.Println(scanner.Text())
		// 한 줄씩 출력
		if err != nil{
		log.Fatal(err)
	}
		i++
	}

	err = data.Close()
	if err != nil{
		log.Fatal(err)
	}
	if scanner.Err() != nil{
		log.Fatal(scanner.Err())
	}

	fmt.Println(numbers)
}
// 현재 프로그램의 문제점
// 파일에 숫자가 3개보다 적거나 많으면 에러 발생 -> 유연하지 못함
// 해결방안: 슬라이스 사용
