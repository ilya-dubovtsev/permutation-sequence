package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	rest := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		rest = append(rest, strconv.Itoa(i))
	}

	return getPermutationByIndex(k, rest)
}

func getPermutationByIndex(index int, rest []string) string {
	if len(rest) == 1 {
		return rest[0]
	}
	if index <= 2 && len(rest) == 2 {
		if index == 1 {
			return strings.Join(rest, "")
		} else {
			return rest[1] + rest[0]
		}
	}

	f := Factorial(len(rest) - 1)

	grade := (index - 1) / f
	first := rest[grade]
	rest = append(rest[:grade], rest[grade+1:]...)

	return first + getPermutationByIndex(index-(grade*f), rest)
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {

	//Input: n = 3, k = 3
	//Output: "213"
	fmt.Println(getPermutation(3, 3))

	//Input: n = 4, k = 9
	//Output: "2314"
	fmt.Println(getPermutation(4, 9))

	// "1"
	fmt.Println(getPermutation(1, 1))
	// "132"
	fmt.Println(getPermutation(3, 2))
	// "1432"
	fmt.Println(getPermutation(4, 6))
}
