package problem_7

import "github.com/kellenff/euler-go/util"

func nthPrime(n uint) uint64 {
	count := uint(0)
	num := uint64(2)

	for count < n {
		if util.IsPrime(num) {
			count++
		}
		num++
	}

	return num - 1
}

func Solve() uint64 {
	return nthPrime(10001)
}
