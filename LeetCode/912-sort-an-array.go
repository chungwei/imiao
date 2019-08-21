package main

import (
	"fmt"
)

/**
912. 排序数组

给定一个整数数组 nums，将该数组升序排列。

示例 1：
输入：[5,2,3,1]
输出：[1,2,3,5]

示例 2：
输入：[5,1,1,2,0,0]
输出：[0,0,1,1,2,5]

提示：
1 <= A.length <= 10000
-50000 <= A[i] <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	n := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(bSort(n))
	n = []int{5, 1, 1, 2, 0, 0}
	fmt.Println(qSort(n, 0, len(n)-1))

	n = []int{5, 2, 3, 1}
	fmt.Println(bSort(n))
	n = []int{5, 2, 3, 1}
	fmt.Println(qSort(n, 0, len(n)-1))

	n = []int{3, 6, 1, 4, 2, 8}
	fmt.Println(bSort(n))
	n = []int{3, 6, 1, 4, 2, 8}
	fmt.Println(qSort(n, 0, len(n)-1))

	n = []int{7, 1, 3, 9, 6, 5, 8}
	fmt.Println(bSort(n))
	n = []int{7, 1, 3, 9, 6, 5, 8}
	fmt.Println(qSort(n, 0, len(n)-1))

}

// 冒泡
func bSort(nums []int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		f := false
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				t := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = t
				f = true
			}
		}
		if !f {
			break
		}
	}
	return nums
}

/**
快速排序
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]；
3）从j开始向前搜索，即由后开始向前搜索(j--)，找到第一个小于key的值A[j]，将A[j]和A[i]的值交换；
4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]的值交换；
5）重复第3、4步，直到i=j；
(3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，使得j=j-1，i=i+1，直至找到为止。
找到符合条件的值，进行交换的时候i， j指针位置不变。
另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。
*/
func qSort(nums []int, left, right int) []int {
	if left >= right {
		return nums
	}
	// 将大于k的都放在右边，小于的，都放在左边
	i, j, k := left, right, nums[left]
	for i < j {
		// 如果flag从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比flag小的
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
