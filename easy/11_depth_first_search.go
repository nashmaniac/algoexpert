package easy

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	if len(n.Children) == 0 {
		return array
	}
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
	return array
}
