package main

import "fmt"

func calc(x []int64) []int64 {
	i := x[len(x)-2]
	for i < x[len(x)-1] {
		i++
		if x[len(x)-1]%i == 0 && (x[len(x)-1] != i) {
			var new_x = make([]int64, len(x)+1)
			copy(new_x, x)
			new_x[len(new_x)-2] = i
			new_x[len(new_x)-1] = x[len(x)-1] / i
			return calc(new_x)
		}
	}
	return x
}

func main() {
	primes := []int64{1, 600851475143}
	primes = calc(primes)
	fmt.Printf("ans: %d\n", primes[len(primes)-1])
}
