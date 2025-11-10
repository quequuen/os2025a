// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0 //그냥 0이라고 쓰면 타입 추론으로 float64가 아닌 int로 잡힘. 그래서 에러!
	// for _, number := range numbers {
	for i := 0; i < len(numbers); i++ {
		fmt.Println("number",i,"의 값: ",numbers[i])
		sum = sum + numbers[i]

	}
	weeks := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/weeks)

}
