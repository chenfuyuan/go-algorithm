package is_valid_bst

import "go-algorithm/src/main/datastruct"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

var Pre int

// IsValidBSTLdrRecursion
// 二叉搜索树 中序遍历后，序列有序
// 则使用中序遍历后，发现序列无序，则表示非二叉搜索树
// 但是存在问题，每次都需要对Pre参数进行初始化为最小值。不利于并发
func IsValidBSTLdrRecursion(root *datastruct.BitTreeNode) bool {
	Pre = INT_MIN
	return helper(root)
}

func helper(root *datastruct.BitTreeNode) bool {
	if root == nil {
		return true
	}
	if !helper(root.Left) {
		return false
	}
	if root.Val <= Pre {
		//此处要<= ，因为可能存在前一个等于 当前值的情况，这种也不为二叉搜索树
		return false
	}
	Pre = root.Val
	if !helper(root.Right) {
		return false
	}
	return true
}
