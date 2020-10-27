package main

import (
	"fmt"
	"strings"
)

func sort(s string, indices []int) string {
	runes := []rune(s)
	sorted := make([]string, len(runes))

	for i := range runes {
		sorted[indices[i]] = string(runes[i])
	}

	return strings.Join(sorted, "")
}

func main() {
	cases := map[string][]int{
		"aiohn": {3, 1, 4, 2, 0},
		"art":   {1, 0, 2},
		"园工室作家": {1, 2, 4, 3, 0},
	}

	for s, indices := range cases {
		fmt.Println(s, "=>", sort(s, indices))
	}
}
