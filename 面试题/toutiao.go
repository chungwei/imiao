package main

import (
	"fmt"
)

func main() {
	n := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(qSort(n, 0, len(n)-1))
	n = []int{5, 8, 1, 2, 0, 40}
	fmt.Println(qSort(n, 0, len(n)-1))

	n = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(rain(n))
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
	l := len(height)
	if l < 3 {
		return 0
	}

	sum, max := 0, 0
	for k, v := range height {
		if height[max] < v {
			max = k
		}
	}

	t := 0
	for i := 0; i <= max; i++ {
		if t <= height[i] {
			t = height[i]
		} else {
			sum += t - height[i]
		}
	}

	t = 0
	for i := l - 1; i >= max; i-- {
		if t <= height[i] {
			t = height[i]
		} else {
			sum += t - height[i]
		}
	}

	return sum
}

func candy(ratings []int) int {
	l := len(ratings)

	return 0
}
