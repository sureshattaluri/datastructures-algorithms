package datastructures

func MinHeightBST(array []int) *BST {
	return buildBST(array, 0, len(array)-1)
}

func buildBST(array []int, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}
	midIdx := (startIdx + endIdx)/2
	bst := &BST{Value: array[midIdx]}
	bst.Left = buildBST(array, startIdx, midIdx-1)
	bst.Right = buildBST(array, midIdx+1, endIdx)
	return bst
}
