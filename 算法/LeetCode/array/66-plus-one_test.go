package array

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, []int{1, 7, 3, 6, 5, 7}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{0}, []int{1}},
		{[]int{1, 2, 3, 9}, []int{1, 2, 4, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
	}

	for _, v := range tests {
		actual := plusOne(v.in)
		if ok := reflect.DeepEqual(v.expected, actual); !ok {
			t.Error("fail in:", v.in, "||expected:", v.expected, "||actual:", actual)
		} else {
			t.Log("done in:", v.in, "||expected:", v.expected, "||actual:", actual)
		}
	}
}
