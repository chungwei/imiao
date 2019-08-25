package main

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

	//root := new(TreeNode)
	//root.Val = 1
	//
	//root.Left = nil
	//
	//node := new(TreeNode)
	//node.Val = 2
	//node.Right = nil
	//
	//root.Right = node
	//
	//node = new(TreeNode)
	//node.Val = 3
	//node.Left = nil
	//node.Right = nil
	//
	//root.Right.Left = node

}

// 后序遍历递归实现
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	l := postorderTraversal(root.Left)
	ret = append(ret, l...)

	r := postorderTraversal(root.Right)
	ret = append(ret, r...)

	ret = append(ret, root.Val)

	return ret
}

// 后序遍历非递归实现
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ret []int
	var stack []TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			t := []int{root.Val}
			ret = append(t, ret...) // 注意：每次新加的元素都是放在首位

			stack = append(stack, *root) // 入栈
			root = root.Right            // 这里跟先序和中序不同
		} else {
			root = &stack[len(stack)-1] // 出栈
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}

	return ret
}
