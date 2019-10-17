package array

func dominantIndex(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	max, second := 0, 0
	for i := 1; i < l; i++ {
		if nums[max] < nums[i] {
			second = nums[max]
			max = i
			continue
		}
		if second < nums[i] {
			second = nums[i]
		}
	}
	if nums[max] >= 2*second {
		return max
	}

	return -1
}
