package db

import "fmt"

/**
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	x := []int{3, 2, 3}
	fmt.Println(majorityElement(x))

	x = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(x))
}

func majorityElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	mid := int(l / 2)
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}

		if m[v] > mid {
			return v
		}
	}
	return 0
}
