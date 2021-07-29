package hasduperune

import "fmt"

func ExampleHasDupeRune() {
	fmt.Println(HasDupeRune("숨바꼭질"))
	fmt.Println(HasDupeRune("다시합시다"))

	// Output:
	// false
	// true
}
