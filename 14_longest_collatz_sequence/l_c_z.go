package main

import "fmt"

func collatz(num int64, known_res *map[int64]int64) (res int64) {
	if value, ok := (*known_res)[num]; ok {
		res = value
	} else if num == 1 {
		res = 1
	} else if num%2 == 0 {
		res = collatz(num/2, known_res) + 1
	} else {
		res = collatz(3*num+1, known_res) + 1
	}
	(*known_res)[num] = res
	return
}

func main() {
	res_set := make(map[int64]int64, 1000000)
	for i := int64(1); i <= 1000000; i++ {
		collatz(i, &res_set)
	}
	res := int64(1)
	for key, value := range res_set {
		if value > res_set[res] {
			res = key
		} else {
			continue
		}
	}
	fmt.Printf("Longest_Collatz_Squence: Collatz(%d)=%d\n", res, res_set[res])
}
