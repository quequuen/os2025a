package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
	//외부 파일을 쓸 시에도 go get 해당 주소 명령어로 다운을 받아야 한다.
	// "week10/pkg/utils/keyboard"
)

func main() {
	fmt.Print("점수 입력 : ")
	n, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.1f\n",n)
}
