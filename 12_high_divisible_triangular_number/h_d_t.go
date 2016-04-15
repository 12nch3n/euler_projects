package main

import (
	"enriqueChen/euler_projects/utils"
	"fmt"
)

func main() {
	var num int64 = 0
	for i := int64(1); ; i++ {
		num = num + i
		divisors := utils.Divisors(num)
		if divisors.Len() > 500 {
			break
		}
	}
	fmt.Printf("ans = %d\n", num)
}
