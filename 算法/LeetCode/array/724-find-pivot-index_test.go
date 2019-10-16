package array

import (
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{-1, -1, -1, 0, 1, 1}, 0},
		{[]int{10}, -1},
		{[]int{}, -1},
		{[]int{-1, -1, 0, 1, 1, 0}, 5},
		{[]int{0, 0, 0}, 0},
		{[]int{0}, -1},
		{[]int{0, 1}, 1},
		{[]int{1, 0}, 0},
	}

	for _, v := range tests {
		actual := pivotIndex(v.in)
		t.Log("in:", v.in, "||expected:", v.expected, "||actual:", actual)
		if actual != v.expected {
			t.Error("fail")
		}
	}
}
