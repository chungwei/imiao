package main

import (
	"fmt"
	"math/rand"
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

解法参考文献 https://www.cnblogs.com/onepixel/p/7674659.html

*/
func main() {
	size := 6
	n, n1 := initInput(size)
	fmt.Println(`冒泡:`, n, `=>`, bubbleSort(n1))

	n, n1 = initInput(size)
	fmt.Println(`快排:`, n, `=>`, quickSort(n1, 0, len(n1)-1))

	n, n1 = initInput(size)
	fmt.Println(`归并:`, n, `=>`, mergeSort(n1))

	n, n1 = initInput(size)
	fmt.Println(`选择:`, n, `=>`, selectSort(n1))

	n, n1 = initInput(size)
	fmt.Println(`插入:`, n, `=>`, insertSort(n1))

	n, n1 = initInput(size)
	fmt.Println(`计数排序:`, n, `=>`, countingSort(n1))

}

// 初始化待排序的数组
func initInput(n int) ([]int, []int) {
	r := []int{}
	for i := 0; i < n; i++ {
		r = append(r, rand.Intn(10))
	}
	r1 := make([]int, n)
	copy(r1, r)
	return r, r1
}

// 冒泡
func bubbleSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}

	for i := 0; i < l-1; i++ {
		f := true
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				f = false
			}
		}
		if f {
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
func quickSort(n []int, left, right int) []int {
	if left > right {
		return n
	}

	k, i, j := n[left], left, right
	for i < j {
		for i < j && n[j] >= k {
			j--
		}
		n[i] = n[j]

		for i < j && n[i] <= k {
			i++
		}
		n[j] = n[i]
	}
	n[i] = k
	quickSort(n, left, i-1)
	quickSort(n, i+1, right)
	return n
}

// 选择排序
func selectSort(n []int) []int {
	l := len(n)
	if l <= 1 {
		return n
	}

	for i := 0; i < l; i++ {
		idx := i
		for j := i; j < l; j++ {
			if n[idx] > n[j] { // 找到最小
				idx = j
			}
		}
		if idx != i { // 把最小值放到已排序列末端
			t := n[i]
			n[i] = n[idx]
			n[idx] = t
		}
	}
	return n
}

// 归并
func mergeSort(n []int) []int {
	l := len(n)
	if l <= 1 {
		return n
	}
	mid := l / 2
	left := mergeSort(n[:mid])
	right := mergeSort(n[mid:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	i, j, ll, rl := 0, 0, len(left), len(right)
	for i < ll && j < rl {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return
}

// 插入排序
func insertSort(n []int) []int {
	l := len(n)
	if l <= 1 {
		return n
	}

	for i := 1; i < l; i++ {
		t, j := n[i], i-1
		for j >= 0 && t <= n[j] {
			n[j+1] = n[j]
			j--
		}
		n[j+1] = t
	}

	return n
}

func countingSort(n []int) []int {
	l := len(n)
	if l <= 1 {
		return n
	}
	max := n[0]
	for _, v := range n {
		if max < v {
			max = v
		}
	}
	sortedArr := make([]int, len(n))
	countsArr := make([]int, max+1) // max+1 是为了防止 countsArr[] 计数时溢出

	// 元素计数
	for _, v := range n {
		countsArr[v]++
	}

	// 统计独特数字个数并累加
	for i := 1; i <= max; i++ {
		countsArr[i] += countsArr[i-1]
	}

	// 让 arr 中每个元素找到其位置
	for _, v := range n {
		sortedArr[countsArr[v]-1] = v
		//fmt.Print(countsArr[v]-1, " ")
		// 保证稳定性
		countsArr[v]--
	}

	return sortedArr
}
