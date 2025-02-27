package algorithms

func LinearSearch(arr []int, v int) bool {
	for _, val := range arr {
		if val == v {
			return true
		}
	}
	return false
}
