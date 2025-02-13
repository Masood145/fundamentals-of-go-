package main

import (
	"fmt"
	"fundamentals-of-go-/learning"
)

func main() {

	/*	learning.Const()
		learning.Variable()
		learning.Values()
		learning.Forloop()
		learning.Ifelse()
		learning.Switch()
		learning.Array()
		learning.Slice()
		learning.Map()
	*/
	/*
		fmt.Println(learning.Add(2, 4))
		fmt.Println(learning.Addm(2, 4, 9))
		a, b := learning.Vals()
		fmt.Println(a, b)
		_, c := learning.Vals()
		fmt.Println(c)
	*/

	/*
		learning.Varfunc(2, 4)
		learning.Varfunc(2, 3, 4, 5)

		nums := []int{1, 2, 3, 4}
		learning.Varfunc(nums...)
	*/

	/*
		va := learning.Clouser()
		fmt.Println(va())
		fmt.Println(va())
		fmt.Println(va())
		va1 := learning.Clouser()
		fmt.Println(va1())
	*/
	/*
		fmt.Println(learning.Recur(7))
		var fib func(n int) int

		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)

		}
		fmt.Println(fib(7))
	*/

	//learning.Rangeovertypes()

	i := 1
	learning.Zeroval(i)
	fmt.Println(i)
	learning.Zeropoint(&i)
	fmt.Println(i)
	fmt.Println(&i)
}
