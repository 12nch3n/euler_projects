package main

import (
	"fmt"
	"math"
)

func IsPalind(num int) bool {
	num_len := int(math.Floor(math.Log10(float64(num))) + 1)
	for i := 1; i <= num_len/2+1; i++ {
		a := num / int(math.Pow(10, float64(i-1))) % 10
		b := num / int(math.Pow(10, float64(num_len-i))) % 10
		if a != b {
			return false
		}
	}
	return true
}

func main() {
	IsPalind(12321)
	for p := 1; p <= 9; p++ {
		n := p*9 + 1
		for j := 1; j <= n; j++ {
			if n%j == 0 {
				a := 1001 - n/j
				b := 1001 - 11*j
				ans := a * b
				if IsPalind(ans) {
					fmt.Printf("ans: %d\n", ans)
				}
			}
		}
	}
}
