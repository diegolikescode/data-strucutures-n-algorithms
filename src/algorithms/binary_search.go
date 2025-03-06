package algorithms

func BinarySearch(list []int, f int) bool {
	lo := 0
	hi := len(list)

	for lo < hi {
		n := lo + (hi-lo)/2
		if list[n] == f {
			return true
		} else if f > list[n] {
			lo = n + 1
		} else if f < list[n] {
			hi = n
		}

	}

	return false
}
