package tree

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

// 中序遍历递归实现
func inorderTraversal(root *TreeNode) []int {
	ret := new([]int)
	InOT(root, ret)
	return *ret
}

func InOT(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	InOT(root.Left, ret)          // 优先递归访问左子树
	*ret = append(*ret, root.Val) // 输出根结点
	InOT(root.Right, ret)         // 递归访问左子树
}

// 中序遍历非递归实现
func InOTR(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	var stack []TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, *root) // 根结点入栈
			root = root.Left             // 优先访问左子树
		} else {
			root = &stack[len(stack)-1] // 说明左子树为空,则根结点出栈
			stack = stack[:len(stack)-1]
			*ret = append(*ret, root.Val) // 输出根结点,即栈顶元素
			root = root.Right
		}
	}
}
