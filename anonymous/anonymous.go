package anonymous

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	tripleFunc := createTransformFunc(3)
	tNumbers := transformNumbers(&numbers, tripleFunc)
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Doubled numbers:", dNumbers)
	fmt.Println("Tripled numbers:", tNumbers)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, num := range *numbers {
		dNumbers = append(dNumbers, transform(num))
	}
	return dNumbers
}

func createTransformFunc(multiplier int) transformFunc {
	return func(number int) int {
		return number * multiplier
	}
}
