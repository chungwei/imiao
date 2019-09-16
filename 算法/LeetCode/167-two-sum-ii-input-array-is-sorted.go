package db

import "fmt"

func main() {
	n := []int{1, 2, 3, 5, 6, 99}
	fmt.Println(twoSumII(n, 9))
}
func twoSumII(numbers []int, target int) []int {
	l, ret := len(numbers), []int{}
	if l < 2 {
		return ret
	}
	m := make(map[int]int, l)
	for k, v := range numbers {
		m[v] = k + 1
	}
	for k, v := range numbers {
		t := target - v
		if i, ok := m[t]; ok && (k+1) != i {
			if (k + 1) > i {
				return []int{i, k + 1}
			}
			return []int{k + 1, i}
		}
	}
	return ret

}
