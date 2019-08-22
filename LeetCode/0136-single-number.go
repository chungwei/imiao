package main

/**
136. 只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
输入: [2,2,1]
输出: 1

示例 2:
输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import "fmt"

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))

	nums = []int{4, 1, 2, 2, 1}
	fmt.Println(singleNumber(nums))

	nums = []int{4, 1, 2, 2, 1}
	fmt.Println(singleNumber1(nums))
}

// 去重
func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, i := range nums {
		m[i]++
		if m[i] > 1 { // 说明是重复的，删除元素 最后剩下只有一个元素
			delete(m, i)
		}
	}
	for k := range m {
		return k // 直接返回即可
	}
	return 0
}

// 异或运算，它是满足交换律和结合律，也就是说 a^b^c = a^c^b，这样当我们遍历数组，顺次进行异或运算，那么最终的结果就是唯一的不重复数字。
// 如[4,1,2,1,2]，4^1^2^1^2 = 1^1^2^2^4 = 0^0^4=4
func singleNumber1(nums []int) int {
	r := 0
	for _, v := range nums {
		r = r ^ v
	}

	return r
}
