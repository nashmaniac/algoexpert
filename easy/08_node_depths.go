package easy

func countLevel(root *BinaryTree, currentSum int) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return currentSum
	}

	leftValue := countLevel(root.Left, currentSum+1)
	rightValue := countLevel(root.Right, currentSum+1)
	return currentSum + leftValue + rightValue
}

func NodeDepths(root *BinaryTree) int {
	return countLevel(root, 0)
}
