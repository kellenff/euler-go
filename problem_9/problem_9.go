package problem_9

func Solve() uint64 {
	for a := 0; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			for c := b + 1; c < 1000; c++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					return uint64(a * b * c)
				}
			}
		}
	}
	panic("Unreachable")
}
