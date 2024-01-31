package problem_2

func getFibs(limit uint64) []uint64 {
	fibs := []uint64{1, 2}
	for i := uint64(2); i < limit; i++ {
		val := fibs[i-1] + fibs[i-2]

		if val >= limit {
			break
		}
		fibs = append(fibs, val)
	}
	return fibs
}

func Solve(limit uint64) uint64 {
	fibs := getFibs(limit)
	sum := uint64(0)

	for _, v := range fibs {
		if v%2 == 0 {
			sum += v
		}
	}

	return sum
}
