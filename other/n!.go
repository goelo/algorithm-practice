package main

import "fmt"

func trailingZeroes(n int) int {
	var num int64
	num = 1
	for i := 1; i <= n; i++ {
		num = num * int64(i)
	}
	fmt.Println(num)

	return int(calcDepth(num, 0))
}

func calcDepth(n, depth int64) int64 {
	if n == 0 {
		return depth
	}
	if n%10 != 0 {
		return depth
	}
	depth++
	return calcDepth(n/10, depth)
}
