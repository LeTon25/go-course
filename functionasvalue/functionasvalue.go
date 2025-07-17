package functionasvalue

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	dNumbers := transformNumbers(&numbers, double)
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Doubled numbers:", dNumbers)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, num := range *numbers {
		dNumbers = append(dNumbers, transform(num))
	}
	return dNumbers
}
func getTransformFunction(choice string) transformFunc {
	switch choice {
	case "double":
		return double
	case "triple":
		return triple
	default:
		return nil
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
