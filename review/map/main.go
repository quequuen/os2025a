package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	// 득표수 세는 프로그램

	counts := make(map[string]int)
	data, err := os.Open("./map/vote.txt")
	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewReader(data)

	for {
		line, err := scanner.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		counts[line]++
	}

	err = data.Close()
	if err != nil {
		log.Fatal(err)
	}

	for name, count := range counts{
		fmt.Println(name, ":", count)
	}

	//map 리터럴
	person := map[string]string{
		"name": "Alice",
		"city": "Seoul",
	}
	print(person["name"])
	print(person["city"])

	// 기본 맵 사용법
	myMap := make(map[string]int)
	myMap["key1"] = 10
	myMap["key2"] = 20
}
