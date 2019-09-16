package db

import "fmt"

/**
2. 两数相加

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	type ListNode struct {
		Val  int
		Next *ListNode
	}

	l1 := &ListNode{1, nil}
	//l1.Next = &ListNode{7, nil}
	//l1.Next.Next = &ListNode{3, nil}

	l2 := &ListNode{9, nil}
	l2.Next = &ListNode{9, nil}
	//l2.Next.Next = &ListNode{9, nil}

	//////
	c := 0
	head := new(ListNode)
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
		tmp.Next = &ListNode{sum % 10, nil} // 注意不要使用 (sum+c)%10，否则下面的 c = sum / 10容易出错
		tmp = tmp.Next
		c = sum / 10
	}

	///////
	list := head.Next
	for list != nil {
		fmt.Print(list.Val, `->`)
		list = list.Next
	}
}
