package main

import (
	"enriqueChen/euler_projects/utils"
	"fmt"
)

func main() {
	//var ques, ans int64 = 2000000, 0
	var ques, ans int64 = 10, 0
	primes := utils.FindPrimes(ques)
	for e := primes.Front(); e != nil; e = e.Next() {
		ans = ans + e.Value.(int64)
	}
	fmt.Printf("F(%d) ans = %d\n", ques, ans)
}
