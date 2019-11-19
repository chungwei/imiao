package tree

/**
255. 验证前序遍历序列二叉搜索树

给定一个整数数组，你需要验证它是否是一个二叉搜索树正确的先序遍历序列。
你可以假定该序列中的数都是不相同的。
参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3

示例 1：
输入: [5,2,6,1,3]
输出: false

示例 2：
输入: [5,2,1,3,6]
输出: true

进阶挑战：
您能否使用恒定的空间复杂度来完成此题？
*/

func verifyPreorder(root *TreeNode, arr []int) bool {
	arrLen := len(arr)
	if root == nil && arrLen == 0 {
		return true
	}
	if root == nil || arrLen == 0 {
		return false
	}

	var ret []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}

	if len(ret) != arrLen {
		return false
	}
	for k, v := range ret {
		if arr[k] != v {
			return false
		}
	}

	return true
}
