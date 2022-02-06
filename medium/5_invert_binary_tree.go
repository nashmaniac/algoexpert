package medium

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) Invert() *BinaryTree {
	if tree == nil {
		return nil
	}

	
	rightInvert := tree.Right.Invert()
	leftInvert := tree.Left.Invert()

	tree.Left = rightInvert
	tree.Right = leftInvert

	return tree
}

func (tree *BinaryTree) InvertBinaryTree() {
	tree.Invert()
}
