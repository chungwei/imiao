package main

import (
	"fmt"
)

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(nums[2:2])
	fmt.Println(twoSum1(nums, 10))
	fmt.Println(twoSum2(nums, 9))
}

// O(n^2)
func twoSum1(nums []int, target int) []int {
	cnt := len(nums)
	if cnt == 0 {
		return []int{}
	}

	for i := 0; i < cnt-1; i++ {
		t := target - nums[i]
		for j := i + 1; j < cnt; j++ {
			if t == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 空间换时间
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		m[v] = k
	}
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		if _, ok := m[t]; ok && i != m[t] {
			return []int{i, m[t]}
		}
	}

	return []int{}
}
