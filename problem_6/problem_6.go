package problem_6

func sumOfSquares(start, end int64) (out int64) {
	for i := start; i < end; i++ {
		out += i * i
	}

	return
}

func squareOfSum(start, end int64) (out int64) {
	for i := start; i < end; i++ {
		out += i
	}

	out *= out
	return
}

func Solve() int64 {
	return squareOfSum(1, 101) - sumOfSquares(1, 101)
}
