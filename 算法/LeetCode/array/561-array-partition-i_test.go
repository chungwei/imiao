package array

import "testing"

// go test -v -run Test_arrayPairSum
func Test_arrayPairSum(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, -1},
		{[]int{1, 4, 3, 2}, 4},
		{[]int{1, 1, 0}, -1},
		{[]int{0}, 0},
		{[]int{1}, 0},
		{[]int{1, 0}, 0},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 6, 3, 0}, 1},
		{[]int{1, 0, 3, 6}, 3},
	}

	for _, v := range tests {
		actual := arrayPairSum(v.in)

		if actual != v.expected {
			t.Error("fail in:", v.in, "||expected:", v.expected, "||actual:", actual)
		} else {
			t.Log("done in:", v.in, "||expected:", v.expected, "||actual:", actual)
		}
	}
}
