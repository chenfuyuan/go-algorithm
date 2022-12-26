package is_valid_bst

import (
	"go-algorithm/src/main/datastruct"
	"math"
)

// IsValidBSTLdr 中序遍历，栈方式实现
func IsValidBSTLdr(root *datastruct.BitTreeNode) bool {
	var stack []*datastruct.BitTreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		//找到最左子树
		for root != nil {
			//入栈
			stack = append(stack, root)
			root = root.Left
		}
		//出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//遍历根节点
		if root.Val <= inorder {
			return false
		}
		//替换最小节点
		inorder = root.Val
		//遍历右子树
		root = root.Right
	}
	return true
}
