package array

import (
	"fmt"

	"github.com/hc0517z/practiceGo/hangul"
)

func Example_array() {
	defer fmt.Println("done")
	fruits := [...]string{"사과", "바나나", "토마토", "계란"}
	for _, fruit := range fruits {

		if hangul.HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}
	}

	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 계란은 맛있다.
	// done
}
