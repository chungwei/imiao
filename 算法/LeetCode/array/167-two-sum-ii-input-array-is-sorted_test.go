package array

import "testing"

// go test -v -i -run Test_twoSum
func Test_twoSum(t *testing.T) {
	tests := []struct {
		in       []int
		target   int
		expected []int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 4, []int{1, 3}},
		{[]int{1, 7, 3, 6, 5, 6}, 100, []int{}},
		{[]int{1, 7, 3, 6, 5, 6}, 7, []int{1, 6}},
		{[]int{1}, 7, []int{}},
	}

	for _, v := range tests {
		actual := twoSum(v.in, v.target)

		if !equal(v.expected, actual) {
			t.Error("fail in:", v.in, "||expected:", v.expected, "||actual:", actual)
		} else {
			t.Log("done in:", v.in, "||expected:", v.expected, "||actual:", actual)
		}
	}
}
