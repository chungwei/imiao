package db

import "fmt"

/**
144. 二叉树的前序遍历
94. 二叉树的中序遍历
145. 二叉树的后序遍历
102. 二叉树的层次遍历

给定一个二叉树，返回它的 前序、中序、后序、层序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]

来源：力扣（LeetCode）
链接：
https://leetcode-cn.com/problems/binary-tree-preorder-traversal
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

	root := new(TreeNode)
	root.Val = 1

	root.Left = nil

	node := new(TreeNode)
	node.Val = 2
	node.Right = nil

	root.Right = node

	node = new(TreeNode)
	node.Val = 3
	node.Left = nil
	node.Right = nil

	root.Right.Left = node

	fmt.Println(`二叉树的前序遍历：`)
	fmt.Println(preorderTraversal(root))
	fmt.Println(preorderTraversal1(root))

	n := []int{3, 4, 5}
	n1 := []int{6}
	n = append(n1, n...)
	fmt.Println(n)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历递归实现
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	ret = append(ret, root.Val)

	l := preorderTraversal(root.Left)
	ret = append(ret, l...)

	r := preorderTraversal(root.Right)
	ret = append(ret, r...)

	return ret
}

// 先序遍历非递归实现
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ret []int
	var stack []TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, *root) // 入栈
			root = root.Left
		} else {
			root = &stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 出栈
			root = root.Right
		}
	}
	return ret
}
