package array

import "testing"

// go test -v -run Test_removeElement
func Test_removeElement(t *testing.T) {
	tests := []struct {
		in       []int
		val      int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 4, 6},
		{[]int{1, 7, 3, 6, 5, 6}, 6, 4},
		{[]int{1, 7, 3, 6, 5, 6}, 7, 5},
		{[]int{}, 7, 0},
	}

	for _, v := range tests {
		actual := removeElement(v.in, v.val)

		if actual != v.expected {
			t.Error("fail in:", v.in, "||expected:", v.expected, "||actual:", actual)
		} else {
			t.Log("done in:", v.in, "||expected:", v.expected, "||actual:", actual)
		}
	}
}
