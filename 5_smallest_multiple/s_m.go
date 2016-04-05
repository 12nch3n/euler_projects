package main

import "fmt"

func gcd(x, y int) int {
	var max, min int
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}
	if max%min == 0 {
		return min
	} else {
		return gcd(max%min, min)
	}
}

func main() {
	multiple := 1
	for i := 2; i <= 10; i++ {
		multiple = multiple * i / gcd(multiple, i)
	}
	fmt.Printf("validate F(10)=%d\n", multiple)
	for i := 2; i <= 20; i++ {
		multiple = multiple * i / gcd(multiple, i)
	}
	fmt.Printf("ans=%d\n", multiple)
}
