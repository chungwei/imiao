package tree

/**
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
递归实现
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

/**
非递归实现
*/
func maxDepthByIteration(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode       // 使用slice实现队列
	var level int               // 记录层数
	queue = append(queue, root) // 根结点优先入队列
	for len(queue) > 0 {
		// 当前队列的长度等于该层的节点数
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			root = queue[0] // 出队列
			queue = queue[1:]

			if root.Left != nil {
				queue = append(queue, root.Left) // 左子树入队列
			}
			if root.Right != nil {
				queue = append(queue, root.Right) // 右子树入队列
			}
		}
		level++ // 层数累加
	}
	return level
}
