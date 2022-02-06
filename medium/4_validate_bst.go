package medium

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) isValid(minValue int, maxValue int) bool {
	if tree == nil {
		return true
	}

	if tree.Value < minValue || tree.Value >= maxValue {
		return false
	}

	return tree.Left.isValid(minValue, tree.Value) && tree.Right.isValid(tree.Value, maxValue)
}

func (tree *BST) ValidateBst() bool {
	// Write your code here.
	return tree.isValid(math.MinInt32, math.MaxInt32)
}
