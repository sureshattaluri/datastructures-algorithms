package algorithms


func Fact (n int) int {
	if n <= 2 {
		return n
	}
	return Fact(n-1)
}

