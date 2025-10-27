package main

import "fmt"

func swap(first *int, second *int) {
	temp := 0
	temp = *first
	*first = *second
	*second = temp
	fmt.Printf("%d %d\n", first, second)
}


func main() {

	var a, b int = 10, 20
	fmt.Println(a, b)
	swap(&a,&b)
	fmt.Println(a, b)

	// fmt.Printf("%.2f\n", math.Sqrt(16.0));
	// fmt.Printf("%.2f\n", math.Sqrt(25.0));

	

}