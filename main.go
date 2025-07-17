package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	for key, value := range f {
		fmt.Printf("%s: %.2f\n", key, value)
	}
}

func main() {
	userNames := make([]string, 3)
	userNames = append(userNames, "John")
	userNames = append(userNames, "Alice")

	fmt.Println("User Names:", userNames)

	courseRatings := make(floatMap)

	courseRatings["Go Programming"] = 4.5
	courseRatings["Python Programming"] = 4.7
	courseRatings["JavaScript Programming"] = 4.2

	fmt.Println("Course Ratings:", courseRatings)

}
