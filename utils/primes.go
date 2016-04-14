package utils

import (
	"container/list"
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

func FindPrimes(count int64) *list.List {
	primes := list.New()
	var num int64 = 2
	for int64(primes.Len()) < count {
		if IsPrime(num, primes.Front()) == true {
			primes.PushBack(num)
		}
		num++
	}
	return primes
}
