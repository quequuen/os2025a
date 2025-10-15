package main

import (
	"fmt"
	"reflect"
)
func main(){

	{
		var price int = 100
		fmt.Println("Price is", price, "dollars.")

		var taxRate float64 = 0.08
		var tax float64 = float64(price)*taxRate
		fmt.Println("Tax is", tax, "dollars.")

		var total float64 = float64(price)+tax
		fmt.Println("Total is", total, "dollars.")

		var availableFunds int = 120
		fmt.Println(availableFunds, "dollars available.")
		fmt.Println("Within budget?", total <= float64(availableFunds))
	}
	{
		var length float64 = 1.0
		var width int = 2
		result := length*float64(width)
		fmt.Println(result) 
		fmt.Println(reflect.TypeOf(result))
		//출력은 int로 보이지만, 실제로는 float64
	}
	
}