package problem_1

func Solve(limit uint) uint {
	sum := uint(0)
	i := uint(1)

	for i < limit {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
		i++
	}

	return sum
}
