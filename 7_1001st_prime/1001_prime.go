package main

import (
	"container/list"
	"fmt"
)

func IsPrime(number int, pKnownPrimes *list.Element) bool {
	e := pKnownPrimes
	for e != nil {
		if number%e.Value.(int) == 0 {
			return false
		}
		e = e.Next()
	}
	return true
}

func FindPrimes(count int) *list.List {
	primes := list.New()
	var num int = 2
	for primes.Len() < count {
		if IsPrime(num, primes.Front()) == true {
			primes.PushBack(num)
		}
		num++
	}
	return primes
}

func main() {
	ans := FindPrimes(10001)
	//for e := ans.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}
	fmt.Printf("ans Prime(10001)=%d\n", ans.Back().Value)
}
