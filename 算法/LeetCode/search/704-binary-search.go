package main

import (
	"fmt"
)

/**
704. 二分查找

给定一个n个元素有序的（升序）整型数组 nums 和一个目标值target，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

提示：
你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	n := []int{-10, 2, 3, 9, 10, 11, 21, 44, 100}
	fmt.Println(bSearch(n, 0, len(n)-1, -10))
	fmt.Println(bSearch(n, 0, len(n)-1, 100))
	fmt.Println(bSearch(n, 0, len(n)-1, 21))
	fmt.Println(bSearch(n, 0, len(n)-1, 210))

	fmt.Println(bSearch1(n, 0, len(n)-1, -10))
	fmt.Println(bSearch1(n, 0, len(n)-1, 100))
	fmt.Println(bSearch1(n, 0, len(n)-1, 21))
	fmt.Println(bSearch1(n, 0, len(n)-1, 210))
	fmt.Println(`=============`)

	n = []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(bSearch(n, 0, len(n)-1, 13))
}

// 递归实现
func bSearch(n []int, l, r int, target int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2 // 注意这里不能使用 mid := (r+l)/2
	if n[mid] == target {
		return mid
	} else if n[mid] > target {
		return bSearch(n, l, mid-1, target)
	} else {
		return bSearch(n, mid+1, r, target)
	}

	return -1
}

// 非递归实现
func bSearch1(n []int, l, r int, target int) int {
	if l > r {
		return -1
	}

	for l <= r {
		mid := l + (r-l)/2
		if n[mid] == target {
			return mid
		} else if n[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
