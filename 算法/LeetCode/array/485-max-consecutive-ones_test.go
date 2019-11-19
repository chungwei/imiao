package array

import (
	"testing"
)

// go test -v -run Test_findMaxConsecutiveOnes
func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 1, 0, 0, 1, 0, 0}, 2},
		{[]int{1, 0, 0, 1, 0, 0}, 1},
		{[]int{1, 0, 1, 1, 0, 1, 1, 1, 1, 0}, 4},
		{[]int{0, 0, 1, 1, 0, 1, 1, 1, 1, 0}, 4},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, v := range tests {
		actual := findMaxConsecutiveOnes(v.in)

		if v.expected != actual {
			t.Error("fail in:", v.in, "||expected:", v.expected, "||actual:", actual)
		} else {
			t.Log("done in:", v.in, "||expected:", v.expected, "||actual:", actual)
		}
	}
}
