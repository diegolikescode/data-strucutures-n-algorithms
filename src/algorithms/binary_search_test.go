package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFindNumber(t *testing.T) {
	list := []int{1, 2, 3, 5, 10, 15, 20, 21}

	f := BinarySearch(list, 2)
	assert.Equal(t, f, true, "Should find number in the first half of the sorted array")

	f2 := BinarySearch(list, 20)
	assert.Equal(t, f2, true, "Should find number in the second half of the sorted array")
}

func testFindNumberEdges(t *testing.T) {
	list := []int{-10, -3, -1, 1, 2, 3, 5, 10, 15, 20, 21, 100, 101, 120}

	f := BinarySearch(list, -10)
	assert.Equal(t, f, true, "Should find number at the first position")

	f2 := BinarySearch(list, 120)
	assert.Equal(t, f2, true, "Should find number at last position")
}

func testNotFindNumbers(t *testing.T) {
	list := []int{-10, -3, -1, 1, 2, 3, 5, 10, 15, 20, 21, 100, 101, 120}

	f := BinarySearch(list, -69)
	assert.Equal(t, f, false, "Should NOT find the number (in the beginning)")

	f2 := BinarySearch(list, 4200)
	assert.Equal(t, f2, false, "Should NOT find the number again (in the end)")

	f3 := BinarySearch(list, 4200)
	assert.Equal(t, f3, false, "Should NOT find the number again (in the middle)")
}

func TestBinarySearch(t *testing.T) {
	testFindNumber(t)
	testFindNumberEdges(t)
	testNotFindNumbers(t)
}
