package exercise2

import (
	"fmt"
	"time"
)

func FibonacciImplementByFor(n int) int {
	start := time.Now()

	if n <= 1 {
		return n
	}

	a, b := 0, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	totaltime := time.Since(start)

	fmt.Println("Thoi gian:", totaltime)
	return b
}

func FiboRecursive(n int) int {

	start := time.Now()
	totaltime := time.Since(start)
	var memory = make([]int, n+1)
	calcufibo := fibonacciRecursive(n, memory)
	fmt.Println("Recursive Fibonacci", totaltime)

	return calcufibo

}

func fibonacciRecursive(n int, memory []int) int {
	if memory[n] != 0{
		return memory[n]
	}


	if n == 0 {
		memory[0] = 0
		return 0
	}
	if n == 1 {
		memory[1] = 1
		return 1
	}

	memory[n] = fibonacciRecursive(n-1, memory) + fibonacciRecursive(n-2, memory)

	return memory[n]

}

func fibonacciSlice(n int) int {

	start := time.Now()

	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	totaltime := time.Since(start)
	fmt.Println("Thoi gian:", totaltime)

	return dp[n]
}
