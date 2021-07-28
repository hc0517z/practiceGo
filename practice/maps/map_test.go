package maps

import (
	"fmt"
	"sort"
)

func Example() {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)
	for _, key := range []rune{'가', '나', '다'} {
		fmt.Println(string(key), codeCount[key])
	}

	// Output:
	// 가 1
	// 나 2
	// 다 1
}

func ExampleCount() {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(rune(key)), codeCount[rune(key)])
	}

	// Output:
	// 가 1
	// 나 2
	// 다 1
}
