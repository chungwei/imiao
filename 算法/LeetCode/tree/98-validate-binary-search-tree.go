package tree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Val > root.Left.Val && root.Val < root.Right.Val {
		return true
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}
