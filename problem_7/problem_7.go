package problem_7

func isPrime(n uint64) bool {
	for i := uint64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return n > 1
}

func nthPrime(n uint) uint64 {
	count := uint(0)
	num := uint64(2)

	for count < n {
		if isPrime(num) {
			count++
		}
		num++
	}

	return num - 1
}

func Solve() uint64 {
	return nthPrime(10001)
}
