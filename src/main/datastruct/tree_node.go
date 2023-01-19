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

func (root *BitTreeNode) InorderGenerate() []int {
	nums := []int{}
	var inorder func(node *BitTreeNode)
	inorder = func(node *BitTreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return nums
}
