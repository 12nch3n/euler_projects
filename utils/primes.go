package utils

import (
	"container/list"
	"fmt"
)

func IsPrime(number int64, pKnownPrimes *list.Element) bool {
	e := pKnownPrimes
	for e != nil {
		if number%e.Value.(int64) == 0 {
			return false
		}
		e = e.Next()
	}
	return true
}

func FindPrimes(max int64) *list.List {
	primes := list.New()
	var num int64 = 3
	primes.PushBack(int64(2))
	for num < max {
		if IsPrime(num, primes.Front()) == true {
			fmt.Printf("Prime <%d> Found.\n", num)
			primes.PushBack(num)
		}
		num++
	}
	return primes
}
