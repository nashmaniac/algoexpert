package easy

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func FindBranchSum(root *BinaryTree, currentSum int, result []int) []int {
	if root == nil {
		return result
	}
	if root.Right == nil && root.Left == nil {
		result = append(result, currentSum+root.Value)
		return result
	}

	result = FindBranchSum(root.Left, currentSum+root.Value, result)
	result = FindBranchSum(root.Right, currentSum+root.Value, result)
	return result
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	arr := make([]int, 0)
	arr = FindBranchSum(root, 0, arr)
	return arr
}
