package main

import "fmt"

func main() {
	// 排序
	n := []int{3, 4, 0, 6, 21, 1}
	fmt.Println(bSort(n))
	n = []int{3, 4, 0, 6, 21, 1}
	fmt.Println(qSort(n, 0, len(n)-1))
	fmt.Println(`-----------------------------`)

	// 二分查找
	n = []int{0, 4, 6, 21, 100}
	fmt.Println(bSearch(n, 0, len(n)-1, 0))
	fmt.Println(bSearch(n, 0, len(n)-1, 21))
	fmt.Println(bSearch(n, 0, len(n)-1, 100))
	fmt.Println(bSearch(n, 0, len(n)-1, 1000))
	fmt.Println(bSearch1(n, 0, len(n)-1, 0))
	fmt.Println(bSearch1(n, 0, len(n)-1, 21))
	fmt.Println(bSearch1(n, 0, len(n)-1, 100))
	fmt.Println(bSearch1(n, 0, len(n)-1, 1000))

	n = []int{111}
	fmt.Println(bSearch(n, 0, len(n)-1, 111))
	fmt.Println(bSearch(n, 0, len(n)-1, 21))
	fmt.Println(bSearch1(n, 0, len(n)-1, 111))
	fmt.Println(bSearch1(n, 0, len(n)-1, 21))
	fmt.Println(`-----------------------------`)

	// 接雨水
	n = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(rain(n))
	fmt.Println(rain1(n))
	fmt.Println(`-----------------------------`)

	// 分糖果
	n = []int{1, 2, 2}
	fmt.Println(candy(n))
	n = []int{1, 0, 2}
	fmt.Println(candy(n))
	n = []int{1, 3, 2, 2, 1}
	fmt.Println(candy(n))
	n = []int{1, 2, 87, 87, 87, 2, 1}
	fmt.Println(candy(n))
	n = []int{1, 3, 4, 5, 2}
	fmt.Println(candy(n))
	fmt.Println(`-----------------------------`)

	list := initList(6)
	printList(list)
	list = reverseList(list)
	printList(list)
	list = delNode(list, 3)
	printList(list)
	list = delNode(list, 1)
	printList(list)
	list = delNode(list, 6)
	printList(list)
	fmt.Println(`-----------------------------`)
	fmt.Println(`-----------------------------`)
	fmt.Println(`-----------------------------`)

}

// 冒泡排序
func bSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}

	for i := 0; i < l-1; i++ {
		fl := true
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				fl = false
			}
		}
		if fl {
			break
		}
	}

	return nums
}

// 快排
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
	qSort(nums, i+1, right)
	qSort(nums, left, i-1)
	return nums
}

// 二分查找递归实现
func bSearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if nums[mid] > target {
		return bSearch(nums, left, mid-1, target)
	} else if nums[mid] < target {
		return bSearch(nums, mid+1, right, target)
	}

	return mid
}

// 二分查找非递归实现
func bSearch1(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 接雨水
func rain(height []int) int {
	l := len(height)
	if l < 3 {
		return 0
	}
	max := 0
	for i := 1; i < l; i++ {
		if height[max] < height[i] {
			max = i
		}
	}

	t, sum := 0, 0
	for i := 0; i <= max; i++ {
		if t < height[i] {
			t = height[i]
		} else {
			sum += t - height[i]
		}
	}

	t = 0
	for i := l - 1; i >= max; i-- {
		if t < height[i] {
			t = height[i]
		} else {
			sum += t - height[i]
		}
	}

	return sum
}

func rain1(height []int) int {
	l := len(height)
	if l < 3 {
		return 0
	}

	left, right := 0, l-1
	lt, rt, sum := 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if lt < height[left] {
				lt = height[left]
			} else {
				sum += lt - height[left]
			}
			left++
		} else {
			if rt < height[right] {
				rt = height[right]
			} else {
				sum += rt - height[right]
			}
			right--
		}

	}

	return sum
}

// 分糖果
func candy(ratings []int) int {
	l := len(ratings)
	c, sum := make([]int, l), 0
	for i := 0; i < l; i++ {
		c[i] = 1
	}

	for i := 0; i < l-1; i++ {
		if ratings[i] < ratings[i+1] && c[i+1] <= c[i] {
			c[i+1] = c[i] + 1
		}
	}

	for i := l - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && c[i-1] <= c[i] {
			c[i-1] = c[i] + 1
		}
	}
	for _, v := range c {
		sum += v
	}

	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initList(n int) *ListNode {
	head := new(ListNode)
	tmp := head
	for i := 1; i <= n; i++ {
		node := new(ListNode)
		node.Val = i

		tmp.Next = node
		tmp = node
	}
	return head.Next
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}
	fmt.Println(``)
}

func reverseList(list *ListNode) *ListNode {
	if list == nil || list.Next == nil {
		return list
	}
	var pre, next *ListNode
	for list != nil {
		next = list.Next
		list.Next = pre
		pre = list
		list = next
	}
	return pre
}

func delNode(list *ListNode, n int) *ListNode {
	if list == nil || list.Next == nil {
		return list
	}
	head := list
	for list != nil {
		if n != list.Val {
			list = list.Next
			continue
		}
		if list.Next == nil {
			fmt.Println(&list, list, list.Val)

			list = nil
			fmt.Println(&list, list)

			break
		}
		list.Val = list.Next.Val
		list.Next = list.Next.Next

	}
	return head
}
