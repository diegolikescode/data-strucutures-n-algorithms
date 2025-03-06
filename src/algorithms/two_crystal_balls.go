package algorithms

import "math"

func TwoCrystalBalls(arr []bool) int {
	l := len(arr)
	jmpAmount := math.Floor(math.Sqrt(float64(l)))
	f := false

	last := 0
	curr := 0
	for !f {
		if arr[curr] {
			break
		} else {
			last = curr
			curr += int(jmpAmount)
		}
	}

	for !f {
		if arr[last] {
			return last
		}
		last++
	}

	return 0
}
