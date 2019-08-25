package main

import (
	"fmt"
)

/**
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:
nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0

示例 2:
nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays1(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays1(nums1, nums2))
	/*fmt.Println(lengthOfLongestSubstring1(`bbbbb`))
	fmt.Println(lengthOfLongestSubstring1(`pwwkew`))
	fmt.Println(lengthOfLongestSubstring1(` `))
	fmt.Println(lengthOfLongestSubstring1(`a`))
	fmt.Println(lengthOfLongestSubstring1(`au`))
	fmt.Println(lengthOfLongestSubstring1(`dvdf`))
	fmt.Println(lengthOfLongestSubstring1("asjrgapa"))
	fmt.Println(lengthOfLongestSubstring1("aab"))*/
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1) / 2
	len2 := len(nums2)

}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	max := 0
	cur := 0
	ints := make(map[int32]int)
	for k, v := range s {
		key, ok := ints[v]
		fmt.Println(k, key, ok)

		dis := k - key
		if !ok || dis > cur {
			cur++
		} else {
			cur = dis
		}
		ints[v] = k
		if max < cur {
			max = cur
		}
	}
	return max
}
