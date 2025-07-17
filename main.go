package main

func main() {
	sum := sumup(1, 2, 3, 4, 5)
	println("Sum of 1 to 5 is:", sum)

	numbers := []int{10, 20, 30}
	sum = sumup(numbers...)
	println("Sum of 10, 20, 30 is:", sum)
}

func sumup(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
