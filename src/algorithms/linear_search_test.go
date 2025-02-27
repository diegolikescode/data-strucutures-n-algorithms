package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")

	arr := []int{1, 2, 3}
	r := LinearSearch(arr, 3)

	assert.Equal(t, true, r, "they should be equal")
}
