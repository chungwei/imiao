package db

import "fmt"

/**
19. 删除链表的倒数第N个节点

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func main() {
	l := initList1()
	printList1(l)

	fmt.Println(`--------------`)

	ll := removeNthFromEnd(l, 2)
	printList1(ll)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode1, n int) *ListNode1 {
	return head
}

func initList1() *ListNode1 {
	head := new(ListNode1)
	pl := head

	for i := 1; i < 10; i++ {
		head.Val = i
		head.Next = new(ListNode1)
		head = head.Next
	}
	head.Next = nil

	return pl
}

func printList1(head *ListNode1) {
	for head.Next != nil {
		fmt.Println(head, head.Val)
		head = head.Next
	}
}
