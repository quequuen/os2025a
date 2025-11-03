package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
함수 정의, 포인터 학습
*/

func swap(first *int, second *int) {
	temp := 0
	temp = *first
	*first = *second
	*second = temp
	fmt.Printf("%d %d\n", first, second)
}

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {

	// var a, b int = 10, 20
	// fmt.Println(a, b)
	// swap(&a,&b)
	// fmt.Println(a, b)

	// fmt.Printf("%.2f\n", math.Sqrt(16.0));
	// fmt.Printf("%.2f\n", math.Sqrt(25.0));

	fmt.Print("점수 입력:")
	score, err := GetFloat()

	if err != nil{
		log.Fatal(err)
	}

	// var status string
	status := ""
	if score >=60{
		status = "합격"
	}else{
		status = "불합격"
	}
	fmt.Println("%.2f점은 %s\n", score, status)
	

}