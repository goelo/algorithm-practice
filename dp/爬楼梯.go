package main

import "fmt"

func climbStairs(n int) int {
	// dp[n] = dp[n-1] + dp[n-2]
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
	// dp := make([]int, n)
	// dp[1] = 1
	// dp[2] = 2
	// for i := 3; i < n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n]
}

func main() {
	n := climbStairs(44)
	fmt.Println(n)
}
