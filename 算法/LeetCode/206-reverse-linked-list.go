package main

import (
	"fmt"
)

/**
反转一个单链表。

示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := initList()
	printList(l)

	fmt.Println(`--------------`)

	ll := reverseList(l)
	printList(ll)
}

func initList() *ListNode {
	head := new(ListNode)
	pl := head

	for i := 1; i < 10; i++ {
		pl.Val = i
		pl.Next = new(ListNode)
		pl = pl.Next
	}
	pl.Next = nil

	return head
}

func printList(head *ListNode) {
	for head.Next != nil {
		fmt.Println(head, head.Val)
		head = head.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := new(ListNode)
	t := new(ListNode)
	/**
		Temp = L->Next;
	        L->Next = Prev;
	        Prev = L;
	        L = Temp;
	*/
	for head.Next != nil {
		t = head.Next
		head.Next = pre
		pre = head
		head = t
	}

	return pre
}
