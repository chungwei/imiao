package tree

import "fmt"

/**
102. 二叉树的层次遍历
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func levelOrder(root *TreeNode) {
	ret := new([][]int)
	LvOT(root, 0, ret)
	fmt.Println(`levelOrder =`, ret)

	ret1 := new([][]int)
	LvOTR(root, 0, ret1)
	fmt.Println(`levelOrder =`, ret1)
}

func LvOT(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	var queue []*TreeNode       // 使用slice实现队列
	queue = append(queue, root) // 根结点优先入队列
	for len(queue) > 0 {
		// 当前队列的长度等于该层的节点数
		qlen := len(queue)
		var tmp []int // 暂存每一层的节点元素
		for i := 0; i < qlen; i++ {
			root = queue[0] // 出队列
			queue = queue[1:]

			tmp = append(tmp, root.Val) // 输出出队列的元素
			if root.Left != nil {
				queue = append(queue, root.Left) // 左子树入队列
			}
			if root.Right != nil {
				queue = append(queue, root.Right) // 右子树入队列
			}
		}
		*ret = append(*ret, tmp)
		level++
	}
}

func LvOTR(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	if level < len(*ret) {
		// 说明该层还有节点未输出,输出该节点元素
		(*ret)[level] = append((*ret)[level], root.Val)
	} else {
		// 说明该层无节点输出
		*ret = append(*ret, []int{root.Val})
	}
	LvOTR(root.Left, level+1, ret)  // 递归遍历左子树
	LvOTR(root.Right, level+1, ret) // 递归遍历右子树
}
