package db

/**
给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
找到所有出现两次的元素。
你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

示例：
输入: [4,3,2,7,8,2,3,1]
输出: [2,3]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-duplicates-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 7}
	fmt.Println(findDuplicates(nums))

	nums = []int{1, 3, 4, 5, 3}
	fmt.Println(findDuplicates(nums))

	nums = []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicates(nums))

	nums = []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) []int {
	ret := []int{}
	m := make(map[int]int, len(nums))

	for _, i := range nums {
		m[i]++
		if m[i] > 1 { // 说明是重复的数
			ret = append(ret, i)
		}
	}
	return ret
}
