package problem_10

import "github.com/kellenff/euler-go/util"

func sumOfPrimes(limit uint64) (out uint64) {
	for i := uint64(2); i < limit; i++ {
		if util.IsPrime(i) {
			out += i
		}
	}

	return
}

func Solve() uint64 {
	return sumOfPrimes(2000000)
}
