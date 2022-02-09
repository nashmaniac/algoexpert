package medium

func FindKthLargestValueInBst(tree *BST, k int) int {
	arr := make([]int, 0)
	tree.InOrder(&arr)
	return arr[len(arr)-k]
}
