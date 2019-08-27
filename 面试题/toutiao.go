package main

import "fmt"

func main() {
	n := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(qSort(n, 0, len(n)-1))
	n = []int{5, 8, 1, 2, 0, 40}
	fmt.Println(qSort(n, 0, len(n)-1))
}

func qSort(nums []int, left, right int) []int {
	if left > right {
		return nums
	}

	i, j, k := left, right, nums[left]
	for i < j {
		for i < j && nums[j] >= k {
			j--
		}
		nums[i] = nums[j]

		for i < j && nums[i] <= k {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = k
	qSort(nums, left, i-1)
	qSort(nums, i+1, right)
	return nums
}

func rain(height []int) int {

}

func candy() {

}
