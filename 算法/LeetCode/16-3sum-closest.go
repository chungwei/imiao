package db

/**
16. 最接近的三数之和

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import (
	"fmt"
	"sort"
)

func main() {
	nums, t := []int{-1, 2, 1, -4}, 1
	fmt.Println(threeSumClosest(nums, t))

	nums, t = []int{1, 1, -1, -1, 3}, -1
	fmt.Println(threeSumClosest(nums, t))

	nums, t = []int{1, 1}, -1
	fmt.Println(threeSumClosest(nums, t))

	nums, t = []int{1, 1, -1, -1, 3}, -1
	fmt.Println(threeSumClosest1(nums, t))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	arrLen := len(nums)
	result := nums[0] + nums[1] + nums[arrLen-1]
	var sum int //三个数组相加的值
	if len(nums) < 3 {
		return 0
	}

	for i := 0; i < arrLen; i++ {
		low, high := i+1, arrLen-1
		for low < high {
			sum = nums[i] + nums[low] + nums[high]
			if sum > target {
				high--
			} else {
				low++
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}

	return result

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i

}

func threeSumClosest1(nums []int, target int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}

	// 遍历数组，三三求和，取绝对值最小
	diff, ret := 1<<31-1, 0
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				t := nums[i] + nums[j] + nums[k]
				if t == target { // 相等  说明距离最近
					return t
				}
				if diff > abs(t-target) {
					diff = abs(t - target)
					ret = t
				}
			}
		}
	}
	return ret
}
