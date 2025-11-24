package main

import "fmt"

type Miles float64
type Kilometers float64
type Meter float64

func (km Kilometers) ToMiles() Miles {
	return Miles(km * 0.621371)
}
func (m Meter) ToMiles() Miles {
	return Miles(m * 0.000621371)
}

func main() {
	kmph := Kilometers(150)
	fmt.Printf("%0.2f kilometers equals %0.2f mile per hour\n", kmph, kmph.ToMiles())
	mph := Meter(500)
	fmt.Printf("%0.2f meters equals %0.2f mile per hour\n", mph, mph.ToMiles())
}
