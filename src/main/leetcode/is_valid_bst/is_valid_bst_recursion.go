package is_valid_bst

import (
	"go-algorithm/src/main/datastruct"
	"math"
)

func IsValidBSTRecursion(root *datastruct.BitTreeNode) bool {
	return isValidBSTRecursionHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTRecursionHelper(root *datastruct.BitTreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	val := root.Val
	if val <= min || val >= max {
		return false
	}
	if !isValidBSTRecursionHelper(root.Left, min, val) {
		return false
	}
	if !isValidBSTRecursionHelper(root.Right, val, max) {
		return false
	}
	return true
}
