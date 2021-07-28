package maps

func Count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}
