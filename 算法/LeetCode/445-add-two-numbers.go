package db

import "fmt"

/**
445. 两数相加 II
给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶:
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例:
输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出: 7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

	l1 := &ListNode3{7, nil}
	l1.Next = &ListNode3{2, nil}
	l1.Next.Next = &ListNode3{4, nil}
	l1.Next.Next.Next = &ListNode3{3, nil}

	l2 := &ListNode3{5, nil}
	l2.Next = &ListNode3{6, nil}
	l2.Next.Next = &ListNode3{4, nil}

	//////
	l1 = rvList(l1)
	l2 = rvList(l2)
	c := 0
	head := new(ListNode3)
	tmp := head
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
		tmp.Next = &ListNode3{sum % 10, nil} // 注意不要使用 (sum+c)%10，否则下面的 c = sum / 10容易出错
		tmp = tmp.Next
		c = sum / 10
	}

	///////
	list := head.Next
	list = rvList(list)
	for list != nil {
		fmt.Print(list.Val, `->`)
		list = list.Next
	}
}

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func rvList(list *ListNode3) *ListNode3 {
	if list == nil || list.Next == nil {
		return list
	}

	var pre, next *ListNode3
	for list != nil {
		next = list.Next
		list.Next = pre
		pre = list
		list = next
	}
	return pre
}
