package datastructures


func ReconstructBst(preOrderTraversalValues []int) *BST {
	if len(preOrderTraversalValues) == 0 {
		return nil
	}

	currentValue := preOrderTraversalValues[0]
	rightSubtreeRootIdx := len(preOrderTraversalValues)

	for i := 1;i < len(preOrderTraversalValues); i++ {
		value := preOrderTraversalValues[i]
		if value >= currentValue {
			rightSubtreeRootIdx = i
			break
		}
	}
	return &BST{
		Value: currentValue,
		Left:  ReconstructBst(preOrderTraversalValues[1:rightSubtreeRootIdx]),
		Right: ReconstructBst(preOrderTraversalValues[rightSubtreeRootIdx:]),
	}
}

