package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {
	f1 := 13
	arr1 := []bool{
		false, false, false, false, false, false, false, false, false, false,
		false, false, false, true, true, true, true, true, true, true,
	}

	f2 := 10
	arr2 := []bool{
		false, false, false, false, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true,
	}

	f3 := 9
	arr3 := []bool{false, false, false, false, false, false, false, false, false, true}

	f4 := 0
	arr4 := []bool{true, true, true, true, true, true, true, true, true, true}

	b1 := TwoCrystalBalls(arr1)
	assert.Equal(t, b1, f1, "Should find the correct position where it breaks")

	b2 := TwoCrystalBalls(arr2)
	assert.Equal(t, b2, f2, "Should find the correct position where it breaks (again)")

	b3 := TwoCrystalBalls(arr3)
	assert.Equal(t, b3, f3, "Should find the correct position where it breaks")

	b4 := TwoCrystalBalls(arr4)
	assert.Equal(t, b4, f4, "Should find the correct position where it breaks (again)")
}
