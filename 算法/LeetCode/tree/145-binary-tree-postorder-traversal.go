package tree

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

	fmt.Println(`二叉树的后序遍历：`)
	fmt.Println(postorderTraversal(root))
}

// 后序遍历递归实现
func postorderTraversal(root *TreeNode) []int {
	ret := new([]int)
	PoOTR(root, ret)
	return *ret
}

func PoOTR(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	PoOTR(root.Left, ret)         // 优先递归访问左子树
	PoOTR(root.Right, ret)        // 递归访问右子树
	*ret = append(*ret, root.Val) // 最后输出根结点
}

// 后序遍历非递归实现
func PoOT(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	var stack []TreeNode // 通过slice实现栈
	for root != nil || len(stack) > 0 {
		if root != nil {
			t := []int{root.Val}
			*ret = append(t, *ret...)    // 输出根结点元素,注意:每次新输出的元素都是放在数组首位
			stack = append(stack, *root) // 根结点入栈
			root = root.Right            // 优先访问右子树,这里跟先序和中序不同
		} else {
			root = &stack[len(stack)-1] // 说明右子树为空,则根结点出栈
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}
}
