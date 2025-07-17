package recursion

func main() {
	fact := factorial(5)
	println("Factorial of 5 is:", fact)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
