package recover_binary_search_tree

import "go-algorithm/src/main/datastruct"

// RecoverTree 99. 恢复二叉搜索树
// https://leetcode.cn/problems/recover-binary-search-tree/
func RecoverTree(root *datastruct.BitTreeNode) {
	//中序遍历获取值序列
	//保存中序遍历的值序列
	nums := []int{}
	var inorder func(node *datastruct.BitTreeNode)
	inorder = func(node *datastruct.BitTreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	//找到错误的两个位置i,j
	x, y := findTwoSwapped(nums)
	//重新遍历二叉搜索树，并进行修复
	doRecoverTree(root, 2, x, y)
}

// 进行恢复操作
// root: 二叉树
// count: 剩余交换的次数
// x: 如果等于x则交换为y
// y: 如果等于y则交换为x
func doRecoverTree(root *datastruct.BitTreeNode, count, x, y int) {
	//再次进行一次遍历
	if root == nil {
		return
	}
	val := root.Val
	//交换值
	if val == x || val == y {
		if val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		//记录已交换次数
		count--
		if count == 0 {
			//修复完毕，直接返回，避免后续遍历
			return
		}
	}
	doRecoverTree(root.Left, count, x, y)
	doRecoverTree(root.Right, count, x, y)
}

// findTwoSwapped 查找2个需要交换的值
func findTwoSwapped(nums []int) (int, int) {
	// 在中序遍历后的序列中，存在以下两种情况
	// 1. 错误的两个值不相邻，则存在2个点i,j，有 num[i] > nums[i+1] && nums[j] > nums[j+1]，此时需要交换 nums[i] 和 nums[j+1]的值进行交换
	//    比如{6,3,4,5,2}此时 nums[0](6) > nums[1](3) , nums[3](5) > nums[4](2),即i=0,j=3，交换nums[0]和nums[4]可以保证序列递增
	//2. 错误的两个值相邻，即只存在一点i,有nums[i] > nums[i+1]，则交换nums[i] 和 nums[i+1]即可
	//index1记录第一个错误的点，index2记录第二个错误的点
	index1, index2 := -1, -1
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			index2 = i
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	//x,y获取需要交换的值
	x, y := nums[index1], nums[index2+1]
	return x, y
}
