package problem_4

import (
	"strconv"
)

func reverse(s string) (out string) {
	for _, v := range s {
		out = string(v) + out
	}

	return
}

func isPalindrome(n int64) bool {
	stringified := strconv.FormatInt(n, 10)

	return stringified == reverse(stringified)
}

func Solve() int64 {
	maxPalindrome := int64(0)
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := int64(i * j)
			if isPalindrome(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}
	return maxPalindrome
}
