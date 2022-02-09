package medium

func (tree *BST) InOrder(arr *[]int) {
	if tree == nil {
		return
	}
	tree.Left.InOrder(arr)
	*arr = append(*arr, tree.Value)
	tree.Right.InOrder(arr)
}

func (tree *BST) PreOrder(arr *[]int) {
	if tree == nil {
		return
	}
	*arr = append(*arr, tree.Value)
	tree.Left.PreOrder(arr)
	tree.Right.PreOrder(arr)
}

func (tree *BST) PostOrder(arr *[]int) {
	if tree == nil {
		return
	}
	tree.Left.PostOrder(arr)
	tree.Right.PostOrder(arr)
	*arr = append(*arr, tree.Value)
}

func (tree *BST) InOrderTraverse(array []int) []int {
	arr := make([]int, 0)
	tree.InOrder(&arr)
	return arr
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	arr := make([]int, 0)
	tree.PreOrder(&arr)
	return arr
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	arr := make([]int, 0)
	tree.PostOrder(&arr)
	return arr
}
