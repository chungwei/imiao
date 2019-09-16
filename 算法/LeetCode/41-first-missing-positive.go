package db

import "fmt"

/**
41. 缺失的第一个正数

给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:
输入: [1,2,0]
输出: 3

示例 2:
输入: [3,4,-1,1]
输出: 2

示例 3:
输入: [7,8,9,11,12]
输出: 1

说明:
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {
	n := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(firstMissingPositive(n))

	n = []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(n))

	n = []int{1}
	fmt.Println(firstMissingPositive(n))
}

func firstMissingPositive(nums []int) int {
	l := len(nums)
	if l < 1 {
		return 1
	}
	m := make(map[int]int, l)
	max := 0
	for _, v := range nums {
		m[v] = 1
		if max < v {
			max = v
		}
	}
	max += 2
	for i := 1; i < max; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 0
}
