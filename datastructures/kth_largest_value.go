package datastructures


type treeInfo struct {
	 numberOfNodesVisited int
	 latestVisitedNodeValue int
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	// Write your code here.
	info := &treeInfo{
		numberOfNodesVisited: 0,
		latestVisitedNodeValue: -1,
	}
	ReverseInOrderTraverse(tree, k, info)
	return info.latestVisitedNodeValue
}

func ReverseInOrderTraverse(tree *BST, k int, info *treeInfo) {
	if tree == nil || info.numberOfNodesVisited >= k {
		return
	}

	ReverseInOrderTraverse(tree.Right, k, info)
	if info.numberOfNodesVisited < k {
		info.latestVisitedNodeValue = tree.Value
		info.numberOfNodesVisited += 1
	}
	ReverseInOrderTraverse(tree.Left, k, info)
}
