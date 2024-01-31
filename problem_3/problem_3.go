package problem_3

func getPrimeFactors(n uint64) []uint64 {
	factors := []uint64{}
	for i := uint64(2); i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

func Solve(n uint64) uint64 {
	primeFactors := getPrimeFactors(n)
	largestPrimeFactor := primeFactors[len(primeFactors)-1]
	return largestPrimeFactor
}
