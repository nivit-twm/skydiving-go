package main

import "github.com/pkg/profile"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	defer profile.Start().Stop()

	Fib(30)
}
