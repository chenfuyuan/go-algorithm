package datastruct

type BitTreeNode struct {
	Val   int
	Left  *BitTreeNode
	Right *BitTreeNode
}

func CreateSimple(val int) *BitTreeNode {
	return &BitTreeNode{Val: val}
}

func CreateFull(val int, leftNode *BitTreeNode, rightNode *BitTreeNode) *BitTreeNode {
	return &BitTreeNode{val, leftNode, rightNode}
}

func CreateClone(source *BitTreeNode) *BitTreeNode {
	return &BitTreeNode{source.Val, source.Left, source.Right}
}
