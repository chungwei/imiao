package main

import (
	"fmt"
)

func main() {

	// 二分查找
	n := []int{0, 4, 6, 21, 100}
	fmt.Println(`二分查找：`)
	fmt.Println(n, `0 =>`, bSearch(n, 0, len(n)-1, 0))
	fmt.Println(n, `21 =>`, bSearch(n, 0, len(n)-1, 21))
	fmt.Println(n, `100 =>`, bSearch(n, 0, len(n)-1, 100))
	fmt.Println(n, `1000 =>`, bSearch(n, 0, len(n)-1, 1000))
	fmt.Println(n, `0 =>`, bSearch1(n, 0, len(n)-1, 0))
	fmt.Println(n, `21 =>`, bSearch1(n, 0, len(n)-1, 21))
	fmt.Println(n, `100 =>`, bSearch1(n, 0, len(n)-1, 100))
	fmt.Println(n, `1000 =>`, bSearch1(n, 0, len(n)-1, 1000))

	n = []int{111}
	fmt.Println(n, `111 =>`, bSearch(n, 0, len(n)-1, 111))
	fmt.Println(n, `21 =>`, bSearch(n, 0, len(n)-1, 21))
	fmt.Println(n, `111 =>`, bSearch1(n, 0, len(n)-1, 111))
	fmt.Println(n, `21 =>`, bSearch1(n, 0, len(n)-1, 21))

	n = []int{}
	fmt.Println(n, `21 =>`, bSearch1(n, 0, len(n)-1, 21))
	fmt.Println(`-----------------------------`)
	//
	//// 接雨水
	//fmt.Println(`接雨水：`)
	//n = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//fmt.Println(n, `=>`, rain(n))
	//fmt.Println(n, `=>`, rain1(n))
	//fmt.Println(`-----------------------------`)
	//
	//// 分糖果
	//fmt.Println(`分糖果：`)
	//n = []int{1, 2, 2}
	//fmt.Println(n, `=>`, candy(n))
	//n = []int{1, 0, 2}
	//fmt.Println(n, `=>`, candy(n))
	//n = []int{1, 3, 2, 2, 1}
	//fmt.Println(n, `=>`, candy(n))
	//n = []int{1, 2, 87, 87, 87, 2, 1}
	//fmt.Println(n, `=>`, candy(n))
	//n = []int{1, 3, 4, 5, 2}
	//fmt.Println(n, `=>`, candy(n))
	//fmt.Println(`-----------------------------`)

	// 链表
	//fmt.Print(`链表初始化：`)
	//list := initList(11, 18)
	//printList(list)
	//
	//fmt.Print(`链表反转：`)
	//list1 := reverseList(list)
	//printList(list1)
	//
	//fmt.Println(`链表相加：`)
	//list2 := initList(5, 9)
	//list2.Val = 9
	//printList(list1)
	//printList(list2)
	//list3 := addList2(list1, list2)
	//printList(list3)
	//
	//list = initList(1, 7)
	//fmt.Print(`链表删除中间：`)
	//list = delNode(list, 3)
	//printList(list)
	//
	//fmt.Print(`链表删除头：`)
	//list = delNode(list, 1)
	//printList(list)
	//
	//fmt.Print(`链表删除尾：`)
	//list = delNode(list, 7)
	//printList(list)
	//list = delNode(list, 6)
	//printList(list)
	//
	//list = initList(1, 7)
	//fmt.Print(`链表k个反转：`)
	//list = reverseKGroup(list, 3)
	//printList(list)
	//
	//fmt.Println(`-----------------------------`)

	// 二叉树
	tree := initTree()

	fmt.Print(`层序：`)
	levelOrder(tree)

	fmt.Println(`-----------------------------`)
	//ns := []Node{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 2}, {2, 0}}
	//fmt.Println(cal(ns, 16))

	fmt.Println(`-----------------------------`)

	setPart()
	setPart1()

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
		if nums[mid] == target {
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
	l, sum := len(height), 0
	if l <= 2 {
		return sum
	}

	max, t := 0, 0
	for k, v := range height {
		if height[max] < v {
			max = k
		}
	}

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
	l, sum, lt, rt := len(height), 0, 0, 0
	if l <= 2 {
		return sum
	}
	left, right := 0, l-1
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
func candy(r []int) int {
	sum, l := 0, len(r)
	if l <= 0 {
		return sum
	}
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		ret[i] = 1
	}
	for i := 0; i < l-1; i++ {
		if r[i] < r[i+1] {
			ret[i+1] = ret[i] + 1
		}
	}
	for i := l - 1; i > 0; i-- {
		if r[i-1] > r[i] && ret[i-1] <= ret[i] {
			ret[i-1] = ret[i] + 1
		}
	}
	for i := 0; i < l; i++ {
		sum += ret[i]
	}
	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initList(start, end int) *ListNode {
	head := new(ListNode)
	t := head
	for i := start; i <= end; i++ {
		n := &ListNode{i, nil}
		t.Next = n
		t = t.Next
	}
	return head.Next
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val)
		if list.Next != nil {
			fmt.Print(`->`)
		}
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

	// del head
	if list.Val == n {
		return list.Next
	}

	tmp := list
	for list.Next != nil {
		if list.Next.Val == n {
			if list.Next.Next == nil { // del tail
				list.Next = nil
			} else {
				list.Next = list.Next.Next // del middle
			}
			return tmp
		}
		list = list.Next
	}
	return tmp
}

func addList2(l1, l2 *ListNode) *ListNode {
	c, list := 0, new(ListNode)
	tmp := list
	for l1 != nil || l2 != nil || c > 0 {
		sum := c
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tmp.Next = &ListNode{sum % 10, nil}
		tmp = tmp.Next
		c = sum / 10
	}
	return list.Next
}

/**
用栈，我们把 k 个数压入栈中，然后弹出来的顺序就是翻转的！

这里要注意几个问题：
第一，剩下的链表个数够不够 k 个（因为不够 k 个不用翻转）；
第二，已经翻转的部分要与剩下链表连接起来

作者：powcai
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/kge-yi-zu-fan-zhuan-lian-biao-by-powcai/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}

	list := new(ListNode)
	tmp := list
	for {
		cnt := k
		stack := []*ListNode{}
		t := head
		for cnt > 0 && t != nil {
			stack = append(stack, t)
			t = t.Next
			cnt--
		}
		if cnt > 0 {
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

func levelOrder(root *TreeNode) {
	ret := new([][]int)
	LOTR(root, 0, ret)
	fmt.Println(`levelOrder =`, ret)

	ret1 := new([][]int)
	LOT(root, 0, ret1)
	fmt.Println(`levelOrder =`, ret1)
}

func LOT(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	var queue []*TreeNode       // 使用slice实现队列
	queue = append(queue, root) // 根结点优先入队列
	for len(queue) > 0 {
		// 当前队列的长度等于该层的节点数
		qlen := len(queue)
		var tmp []int // 暂存每一层的节点元素
		for i := 0; i < qlen; i++ {
			root = queue[0] // 出队列
			queue = queue[1:]

			tmp = append(tmp, root.Val) // 输出出队列的元素
			if root.Left != nil {
				queue = append(queue, root.Left) // 左子树入队列
			}
			if root.Right != nil {
				queue = append(queue, root.Right) // 右子树入队列
			}
		}
		*ret = append(*ret, tmp)
		level++
	}
}

func LOTR(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	if level < len(*ret) {
		// 说明该层还有节点未输出,输出该节点元素
		(*ret)[level] = append((*ret)[level], root.Val)
	} else {
		// 说明该层无节点输出
		*ret = append(*ret, []int{root.Val})
	}
	LOTR(root.Left, level+1, ret)  // 递归遍历左子树
	LOTR(root.Right, level+1, ret) // 递归遍历右子树
}

type node struct {
	x float64
	y float64
}

func cal(ns []node, k int) []node {
	if len(ns) < 3 {
		return []node{}
	}
	if k == 1 {
		return ns
	}
	ns = append(ns, ns[0])
	fmt.Println(ns)
	length := getLength(ns)
	kline := length / float64(k)
	var result []node
	var leftlen float64
	var curline = kline
	fmt.Println(kline)
	for i := 1; i < len(ns); i++ {
		len := getLength(ns[i-1:i+1]) - leftlen
		for leftlen+curline <= len && len >= curline {
			data := node{ns[i-1].x + ((leftlen+curline)/len)*(ns[i].x-ns[i-1].x), ns[i].y}
			if ns[i].x == ns[i-1].x {
				data = node{ns[i].x, ns[i-1].y + ((leftlen+curline)/len)*(ns[i].y-ns[i-1].y)}
			}
			result = append(result, data)
			leftlen += curline
			curline = kline
		}
		leftlen -= len
	}
	return result
}

func getLength(ns []node) float64 {
	var length float64
	pre := ns[0]
	for i := 1; i < len(ns); i++ {
		if ns[i].x == pre.x {
			length += abs(ns[i].y - pre.y)
		} else {
			length += abs(ns[i].x - pre.x)
		}
		pre = ns[i]
	}
	fmt.Println(length, ns)
	return length
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func setPart() {
	nums := []int{8, 6, 4, -3, 5, -2, -1, 0, 1, -9, 1, 0}
	l := len(nums)
	if l < 2 {
		return
	}
	left, right, mid := 0, l-1, 0
	for mid <= right {
		tmp := nums[mid]
		if tmp < 0 {
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		} else if tmp == 0 {
			mid++
		} else {
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
		}
	}
	fmt.Println(`setPart >>>>>`, nums)

	return
}

func setPart1() {
	nums := []int{8, 6, 4, -3, 5, -2, -1, 0, 1, -9, 1, 0}
	l := len(nums)
	if l < 2 {
		return
	}
	idx := -1
	for i := 0; i < l; i++ {
		if nums[i] < 0 {
			tmp := nums[i]
			nums[i] = nums[idx+1]
			nums[idx+1] = tmp
			idx++
		}
	}
	fmt.Println(`setPart1>>>>>`, nums)

	return
}
