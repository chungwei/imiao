package main

import "fmt"

/**
162. 寻找峰值
数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞。

示例 1:
输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。

示例 2:
输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。

说明:
你的解法应该是 O(logN) 时间复杂度的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-peak-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	n := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(n))
	fmt.Println(findPeakElement([]int{}))
	fmt.Println(findPeakElement1(n))
}

func findPeakElement(nums []int) int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return l - 1
}

func findPeakElement1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
