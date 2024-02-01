package util

func IsPrime(n uint64) bool {
	for i := uint64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return n > 1
}
