package main

/*
  Reference: https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
  Solution:
	a = k(x^2-y^2)
	b = 2kxy
	c = k(x^2+y^2)
  So if a+b+c=1000 we have: kx(x+y)=500
  set s = x+y, we have kxs=500. (x<s<2x)
*/

import (
	"container/list"
	"enriqueChen/euler_projects/utils"
	"fmt"
)

func main() {
	var ques int64 = 1000
	var divisors *list.List
	var triples *list.List = list.New()
	divisors = utils.Divisors(ques / 2)
	for e1 := divisors.Front(); e1 != nil; e1 = e1.Next() {
		x := e1.Value.(int64)
		for e2 := e1.Next(); (e2 != nil) && (e2.Value.(int64) < 2*x); e2 = e2.Next() {
			s := e2.Value.(int64)
			if 500%(x*s) == 0 {
				k := (500 / x) / s
				a := k * (x*x - (s-x)*(s-x))
				b := 2 * k * x * (s - x)
				c := k * (x*x + (s-x)*(s-x))
				triples.PushBack(a * b * c)
			}
		}
	}
	for z := triples.Front(); z != nil; z = z.Next() {
		fmt.Printf("ans(%d) = %d\n", ques, z.Value.(int64))
	}
}
