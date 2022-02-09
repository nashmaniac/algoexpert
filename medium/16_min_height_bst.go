package medium

func MinHeightBST(array []int) *BST {
	if len(array) == 0 {
		return nil
	}
	if len(array) == 1 {
		return &BST{Value: array[0]}
	}

	// find the mid
	mid := (len(array) - 1) / 2
	node := &BST{Value: array[mid]}
	node.Left = MinHeightBST(array[:mid])
	node.Right = MinHeightBST(array[mid+1:])
	return node
}
