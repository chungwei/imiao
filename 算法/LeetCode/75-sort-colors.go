package db

import "fmt"

func main() {
	s := []int{1, 1, 0, 2, 2, 0}
	sortColors(s)
}

func sortColors(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}

	for i := 0; i < l; i++ {
		a := true
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				t := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = t
				a = false
			}
		}
		if a {
			break
		}
	}
	fmt.Println(nums)
}
