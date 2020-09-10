package merge

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{[]int{260, -330, -3, -133, -598, 471, -750, -367, 86, 270, -428},
			[]int{-750, -598, -428, -367, -330, -133, -3, 86, 260, 270, 471}},
		{[]int{-439, 455, -709, -723, -10, -91, 54, 323, 238, -334, -736},
			[]int{-736, -723, -709, -439, -334, -91, -10, 54, 238, 323, 455}},
	}
	for _, tt := range tests {
		t.Run("testing merge sort", func(t *testing.T) {
			if got := Sort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
