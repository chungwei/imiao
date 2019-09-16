package db

import "fmt"

/**
349. 两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。

示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]

示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	n1 := []int{1, 2, 2, 1}
	n2 := []int{2, 2}
	fmt.Println(intersection(n1, n2))
	fmt.Println(intersection1(n1, n2))

	n1 = []int{4, 9, 5}
	n2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(n1, n2))
	fmt.Println(intersection1(n1, n2))
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		m[v] = k
	}

	ret := []int{}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			delete(m, v)
			ret = append(ret, v)
		}
	}

	return ret
}

func intersection1(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool, len(nums1))
	for _, v := range nums1 {
		m[v] = true
	}

	ret := []int{}
	for _, v := range nums2 {
		if i, ok := m[v]; ok && i {
			ret = append(ret, v)
			m[v] = false
		}
	}

	return ret

}
