package easy

import (
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) ReturnDiffWithNode(target int) (*BST, *int) {
	if tree == nil {
		return nil, nil
	} 

	leftNode, leftDiff := tree.Left.ReturnDiffWithNode(target)
	rightNode, rightDiff := tree.Right.ReturnDiffWithNode(target)

	var node *BST = nil
	var diff *int = nil

	if leftDiff != nil && rightDiff != nil {
		if *leftDiff < *rightDiff {
			node = leftNode
			diff = leftDiff
		} else {
			node = rightNode
			diff = rightDiff
		}
	} else if leftDiff != nil && rightDiff == nil {
		node = leftNode
		diff = leftDiff
	} else if leftDiff == nil && rightDiff != nil {
		node = rightNode
		diff = rightDiff
	}

	currentDiff := int(math.Abs(float64(tree.Value) - float64(target)))

	if diff != nil {
		if *diff < currentDiff {
			return node, diff
		} else {
			return tree, &currentDiff
		}
	} else {
		return tree, &currentDiff
	}

}

func (tree *BST) FindClosestValue(target int) int {
	node, _ := tree.ReturnDiffWithNode(target)
	return node.Value
}
