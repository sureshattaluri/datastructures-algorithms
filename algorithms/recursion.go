package algorithms

import "fmt"

func Fact (n int) int {
	fmt.Printf("\n processing factorial of %v", n)
	if n <= 2 {
		return n
	}
	return n * Fact(n-1)
}

func Reverse (s string) string {
	fmt.Printf("\n processing reverse of %v", s)
	if len(s) == 1 {
		return s
	}
	return string(s[len(s)-1]) + Reverse(s[0:len(s)-1])
}


