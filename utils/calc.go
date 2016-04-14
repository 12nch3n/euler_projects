package utils

import (
	"container/list"
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

func Divisors(num int64) *list.List {
	ret1 := list.New()
	ret2 := list.New()
	for i := int64(1); i*i <= num; i++ {
		if num%i == 0 {
			ret1.PushBack(i)
			if i != num/i {
				ret2.PushFront(num / i)
			}
		}
	}
	ret1.PushBackList(ret2)
	return ret1
}

func MinDivisors(num int64) *list.List {
	var ret = list.New()
	ret.PushBack(num)
	e := ret.Front()
	for e != nil {
		x1, x2 := Divide(e.Value.(int64))
		if x1 != 1 {
			ret.PushBack(x1)
			ret.PushBack(x2)
			e = e.Next()
			ret.Remove(e.Prev())
		} else {
			e = e.Next()
		}
	}
	return ret
}
