package exersizeFibonacci

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	fmt.Println(fib(n-2) + fib(n-1))
	return fib(n-2) + fib(n-1)
}
