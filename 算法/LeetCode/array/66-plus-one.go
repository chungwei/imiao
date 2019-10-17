package array

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return digits
	}

	tmp := make([]int, l)
	copy(tmp, digits)

	i, c := l-1, 0
	tmp[i] = tmp[i] + 1
	for i >= 0 {
		t := tmp[i] + c
		tmp[i] = t % 10
		c = t / 10
		i--
		continue
	}
	if c > 0 {
		return append([]int{1}, tmp...)
	}

	return tmp
}
