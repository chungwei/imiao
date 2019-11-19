package string

import "testing"

// go test -v -run Test_isAnagram
func Test_isAnagram(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected bool
	}{
		{`anagram`, `nagaram`, true},
		{`tar`, `car`, false},
		{``, ``, true},
	}

	for _, v := range tests {
		actual := isAnagram(v.in1, v.in2)

		if actual != v.expected {
			t.Error("in:", v.in1, v.in2, "||expected:", v.expected, "||actual:", actual)
		} else {
			t.Log("in:", v.in1, v.in2, "||expected:", v.expected, "||actual:", actual)
		}
	}
}
