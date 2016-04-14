package test

import (
	"container/list"
	"fmt"
	"testing"

	"enriqueChen/euler_projects/utils"
)

const ARRAY_MAX_CAP int = 101

func TestDivide(t *testing.T) {
	var test_cases = []struct {
		input  int64
		x1, x2 int64
	}{
		{1, 1, 1},
		{2, 1, 2},
		{6, 2, 3},
		{16, 2, 8},
	}
	for _, mt := range test_cases {
		x1, x2 := utils.Divide(mt.input)
		if x1 != mt.x1 || x2 != mt.x2 {
			t.Errorf("Expect Divide %d to (%d, %d), but got (%d, %d)", mt.input, mt.x1, mt.x2, x1, x2)
		}
	}
}

func compare(length int, expected *[ARRAY_MAX_CAP]int64, actual *list.List) bool {
	if length != actual.Len() {
		return false
	} else {
		fmt.Printf("Expected %v\n", expected[:length])
		index := 0
		for e := actual.Front(); e != nil; e = e.Next() {
			fmt.Printf("Compare index(%d): actual = %d \t expected = %d\n",
				index, e.Value.(int64), expected[index])
			if e.Value.(int64) != expected[index] {
				return false
			}
			index++
		}
	}
	return true
}

func TestMinDivisors(t *testing.T) {
	var test_cases = []struct {
		input    int64
		length   int
		expected [ARRAY_MAX_CAP]int64
	}{
		{1, 1, [ARRAY_MAX_CAP]int64{1}},
		{2, 1, [ARRAY_MAX_CAP]int64{2}},
		{4, 2, [ARRAY_MAX_CAP]int64{2, 2}},
		{30, 3, [ARRAY_MAX_CAP]int64{2, 3, 5}},
	}

	for _, mt := range test_cases {
		actual := utils.MinDivisors(mt.input)
		if compare(mt.length, &mt.expected, actual) == false {
			t.Errorf("Dividenumbers of %d is incorrect.", mt.input)
		}
	}
}

func TestDivisors(t *testing.T) {
	var test_cases = []struct {
		input    int64
		length   int
		expected [ARRAY_MAX_CAP]int64
	}{
		{1, 1, [ARRAY_MAX_CAP]int64{1}},
		{2, 2, [ARRAY_MAX_CAP]int64{1, 2}},
		{4, 3, [ARRAY_MAX_CAP]int64{1, 2, 4}},
		{30, 8, [ARRAY_MAX_CAP]int64{1, 2, 3, 5, 6, 10, 15, 30}},
	}

	for _, mt := range test_cases {
		actual := utils.Divisors(mt.input)
		if compare(mt.length, &mt.expected, actual) == false {
			t.Errorf("Dividenumbers of %d is incorrect.", mt.input)
		}
	}
}
