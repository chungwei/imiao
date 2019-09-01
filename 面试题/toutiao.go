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

	// 链表
	fmt.Print(`链表初始化：`)
	list := initList(2, 5)
	printList(list)

	fmt.Print(`链表反转：`)
	list1 := reverseList(list)
	printList(list1)

	fmt.Print(`链表相加：`)
	list2 := initList(5, 9)
	list2.Val = 9
	//list2 = reverseList(list2)
	list3 := addList2(list1, list2)
	printList(list3)

	list = initList(1, 7)
	fmt.Print(`链表删除中间：`)
	list = delNode(list, 3)
	printList(list)

	fmt.Print(`链表删除头：`)
	list = delNode(list, 1)
	printList(list)

	fmt.Print(`链表k个反转：`)
	list = reverseKGroup(list, 3)
	printList(list)

	fmt.Print(`链表删除尾：`)
	list = delNode(list, 7)
	printList(list)

	fmt.Println(`-----------------------------`)

	// 二叉树
	tree := initTree()
	fmt.Print(`先序：`)
	preOrder(tree)
	fmt.Println(``)
	preOrder1(tree)

	fmt.Print(`中序：`)
	inOrder(tree)
	fmt.Println(``)
	inOrder1(tree)

	fmt.Print(`后序：`)
	postOrder(tree)
	fmt.Println(``)
	postOrder1(tree)

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
	qSort(nums, left, i-1)
	qSort(nums, i+1, right)
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
	for k, v := range height {
		if height[max] < v {
			max = k
		}
	}

	t, sum := 0, 0
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

func initList(start, end int) *ListNode {
	head := new(ListNode)
	tmp := head

	for i := start; i <= end; i++ {
		node := &ListNode{i, nil}
		tmp.Next = node
		tmp = node
	}

	return head.Next

}

func printList(list *ListNode) {
	if list == nil {
		fmt.Print(`nil`)
		return
	}

	for list != nil {
		fmt.Print(list.Val, "->")
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

/**
fmt.Print(`链表删除中间：`)
list = delNode(list, 3)
printList(list)

fmt.Print(`链表删除头：`)
list = delNode(list, 6)
printList(list)

fmt.Print(`链表删除尾：`)
list = delNode(list, 1)
printList(list)
*/
func delNode(list *ListNode, n int) *ListNode {
	if list == nil {
		return list
	}
	if list.Val == n { // 删除头
		return list.Next
	}
	head := list
	for list.Next != nil {
		if list.Next.Val == n {
			if list.Next.Next == nil { // 删除尾
				list.Next = nil
				break
			}
			// 删除中间
			list.Next.Val = list.Next.Next.Val
			list.Next.Next = list.Next.Next.Next
			break
		}
		list = list.Next
	}
	return head
}

func addList2(l1, l2 *ListNode) *ListNode {
	fmt.Println(``)
	printList(l1)
	printList(l2)
	carry := 0
	head := new(ListNode)
	node := head
	for l1 != nil || l2 != nil || carry > 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry

		node.Next = &ListNode{sum % 10, nil}
		node = node.Next
		carry = sum / 10
	}

	return head.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	list := new(ListNode)
	tmp := list
	for {
		cnt := k
		stack := []*ListNode{}
		t := head
		for t != nil && cnt > 0 {
			stack = append(stack, t)
			t = t.Next
			cnt--
		}
		if cnt > 0 { // 说明不足k个
			tmp.Next = head
			break
		}
		for len(stack) > 0 {
			tmp.Next = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp = tmp.Next
		}

		tmp.Next = t
		head = t
	}

	return list.Next

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTree() *TreeNode {
	root := new(TreeNode)
	root.Val = 10

	root.Left = &TreeNode{8, nil, nil}
	root.Left.Left = &TreeNode{6, nil, nil}
	root.Left.Right = &TreeNode{9, nil, nil}

	root.Right = &TreeNode{12, nil, nil}
	root.Right.Right = &TreeNode{17, nil, nil}

	return root
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, `->`)
	preOrder(root.Left)
	preOrder(root.Right)
}
func preOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	//stack = append(stack, root)
	for root != nil || len(stack) > 0 {
		if root != nil {
			fmt.Print(root.Val, `->`)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	fmt.Println(``)
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Print(root.Val, `->`)
	inOrder(root.Right)
}

func inOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	//stack = append(stack, root)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			fmt.Print(root.Val, `->`)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	fmt.Println(``)
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Print(root.Val, `->`)
}

func postOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	var list []int
	for root != nil || len(stack) > 0 {
		if root != nil {
			list = append([]int{root.Val}, list...)
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}
	fmt.Println(list)
}

func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}

}

func levelOrder1(root *TreeNode) {
	if root == nil {
		return
	}

}
