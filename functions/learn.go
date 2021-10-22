package functions

import (
	"fmt"
	"strings"
)

func Learn() {
	fmt.Println(add(3,2))
	fmt.Println(convert("John", "Doe"))
	sum(10,20,30)
}

func add(x, y int) int {
	return x + y
}

func convert(firstname, lastname string) (string, string) {
	s1 := strings.ToUpper(firstname)
	return s1, firstname
}
func sum(numbers ...int) {
	total := 0
	for _, v := range numbers {
		total += v
	}

	fmt.Printf("total is %d", total)
}