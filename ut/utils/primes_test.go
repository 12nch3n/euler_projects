package test

import (
	"container/list"
	"enriqueChen/euler_projects/utils"
	"testing"
)

func TestIsPrime(t *testing.T) {
	var test_cases = []struct {
		input    int64
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
	}
	primes := list.New()
	for _, mt := range test_cases {
		actual := utils.IsPrime(mt.input, primes.Front())
		if actual == true {
			primes.PushBack(mt.input)
		}
		if actual != mt.expected {
			t.Errorf("IsPrime %d = %b is wrong!", mt.input, actual)
		}
	}
}
