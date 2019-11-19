package array

import (
	"reflect"
	"testing"
)

func Test_twoSum2(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 9, []int{2, 5}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 10, []int{}},
		{[]int{}, 1, []int{}},
		{[]int{0}, 1, []int{}},
	}

	for _, v := range tests {
		actual := twoSum2(v.nums, v.target)
		if ok := reflect.DeepEqual(v.expected, actual); !ok {
			t.Error("fail in:", v.nums, v.target, "||expected:", v.expected, "||actual:", actual)
		} else {
			t.Log("done in:", v.nums, v.target, "||expected:", v.expected, "||actual:", actual)
		}
	}
}
