package main

import (
	"container/list"
	"fmt"
)

func Divide(x int64) (int64, int64) {
	for i := int64(2); i < x; i++ {
		if x%i == 0 {
			return i, x / i
		} else {
			continue
		}
	}
	return 1, x
}

func main() {
	primes := list.New()
	primes.PushBack(int64(600851475143))
	e := primes.Front()
	for e != nil {
		x1, x2 := Divide(e.Value.(int64))
		if x1 != 1 {
			primes.PushBack(x1)
			primes.PushBack(x2)
			e = e.Next()
			primes.Remove(e.Prev())
		} else {
			fmt.Println(e.Value)
			e = e.Next()
		}
	}
	fmt.Println(primes.Back().Value)
}
