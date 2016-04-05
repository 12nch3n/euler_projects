package main

import (
	"fmt"
)

// F(n) = F(n-1) + F(n-2) = 4F(n-3) + F(n-6)
// X(n) = 4X(n-1) + X(n-2) x0=0 x1=2

func main() {
	var i, j, sum int = 2, 0, 0
	for i < 4e6 {
		sum += i
		var tmp int = i
		i, j = 4*i+j, tmp
		fmt.Printf("i = %d\tj=%d\n", i, j)
	}
	fmt.Println(sum)
}
