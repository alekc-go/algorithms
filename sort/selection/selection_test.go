package selection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	data := []int{3, 5, 8, 4, 1, 9, -2}
	assert.Equal(t, []int{-2, 1, 3, 4, 5, 8, 9}, Sort(data))
}
