package main

import "fmt"

/**
101. 对称二叉树

给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:
如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode101 struct {
	Val   int
	Left  *TreeNode101
	Right *TreeNode101
}

func main() {
	t1 := &TreeNode101{}
	fmt.Println(isSymmetric(t1))

	t1 = &TreeNode101{1, nil, nil}
	t1.Left = &TreeNode101{2, nil, nil}
	t1.Right = &TreeNode101{2, nil, nil}

	t1.Left.Left = &TreeNode101{3, nil, nil}
	t1.Left.Right = &TreeNode101{4, nil, nil}

	t1.Right.Left = &TreeNode101{4, nil, nil}
	t1.Right.Right = &TreeNode101{3, nil, nil}
	fmt.Println(isSymmetric(t1))

	t1 = &TreeNode101{1, nil, nil}
	fmt.Println(isSymmetric(t1))

	t1 = &TreeNode101{1, nil, nil}
	t1.Left = &TreeNode101{2, nil, nil}
	t1.Right = &TreeNode101{2, nil, nil}

	t1.Left.Left = &TreeNode101{3, nil, nil}

	t1.Right.Left = &TreeNode101{4, nil, nil}
	t1.Right.Right = &TreeNode101{3, nil, nil}
	fmt.Println(isSymmetric(t1))
}
func isSymmetric(root *TreeNode101) bool {
	if root == nil {
		return true
	}
	return checkLR(root.Left, root.Right)
}

func checkLR(l, r *TreeNode101) bool {
	if l == nil || r == nil {
		if l == nil && r == nil {
			return true
		}
		return false
	}

	if l.Val == r.Val {
		return checkLR(l.Left, r.Right) && checkLR(l.Right, r.Left)
	}

	return false
}
