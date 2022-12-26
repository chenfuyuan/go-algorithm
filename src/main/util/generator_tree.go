package util

import "go-algorithm/src/main/datastruct"

func GeneratorTree(source []int) *datastruct.BitTreeNode {
	if source == nil {
		return nil
	}
	return buildTreeNode(source, 0)
}

func buildLeftTreeNode(source []int, rootIndex int) *datastruct.BitTreeNode {
	leftNodeIndex := rootIndex*2 + 1
	return buildTreeNode(source, leftNodeIndex)
}

func buildRightTreeNode(source []int, rootIndex int) *datastruct.BitTreeNode {
	rightNodeIndex := rootIndex*2 + 2
	return buildTreeNode(source, rightNodeIndex)
}

func buildTreeNode(source []int, treeNodeIndex int) *datastruct.BitTreeNode {
	length := len(source)
	if treeNodeIndex >= length {
		return nil
	}
	val := source[treeNodeIndex]
	if val == -1 {
		return nil
	}
	treeNode := datastruct.CreateSimple(val)
	//构建左子树
	treeNode.Left = buildLeftTreeNode(source, treeNodeIndex)
	//构建右子树
	treeNode.Right = buildRightTreeNode(source, treeNodeIndex)
	return treeNode
}
