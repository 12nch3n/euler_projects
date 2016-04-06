package main

import (
	"fmt"
)

func Calc(length float64) float64 {
	var sum float64 = 0
	for i := 1; i <= int(length); i++ {
		sum += ((1+length)/2*length - float64(i)) * float64(i)
	}
	return sum
}

func Calc_simple(length float64) float64 {
	var sum1, sum2 float64 = 0, 0
	for i := 1; i <= int(length); i++ {
		sum1 += float64(i) * float64(i)
		sum2 += float64(i)
	}
	return sum2*sum2 - sum1
}
func main() {
	fmt.Printf("Validate F(10)=%v\n", Calc(10))
	fmt.Printf("Validate F(100)=%v\n", int64(Calc(100)))
}
