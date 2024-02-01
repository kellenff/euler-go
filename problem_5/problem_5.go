package problem_5

func divisibleByRange(min, max int64) int64 {
Outer:
	for i := int64(max); true; i++ {
		for j := max; j > min; j-- {
			if i%j != 0 {
				continue Outer
			}
		}
		return i
	}
	panic("Unreachable")
}

func Solve() int64 {
	return divisibleByRange(1, 20)
}
