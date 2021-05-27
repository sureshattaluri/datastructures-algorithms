package datastructures

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	currentNode := tree
	for {
		if value < currentNode.Value {
			//left
			if currentNode.Left == nil {
				currentNode.Left = &BST{Value: value}
				break
			}
			currentNode = currentNode.Left
		} else {
			if currentNode.Right == nil {
				currentNode.Right = &BST{Value: value}
				break
			}
			currentNode = currentNode.Right
		}
	}
	return tree
}


func (tree *BST) Contains(value int) bool {
	currentNode := tree
	for currentNode != nil {
		if currentNode.Value == value {
			return true
		}
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return false
}

func (tree *BST) Remove(value int) *BST {
	tree.remove(value, nil)
	return tree
}

func (tree *BST) remove(in int, parent *BST) {
	current := tree
	for current != nil {
		if in < current.Value {
			parent = current
			current = current.Left
		} else if in > current.Value {
			parent = current
			current = current.Right
		} else {
			if current.Left != nil && current.Right != nil {
				current.Value = current.Right.getMinValue()
				current.Right.remove(current.Value, current)
			} else if parent == nil {
				if current.Left != nil {
					current.Value = current.Left.Value
					current.Right = current.Left.Right
					current.Left = current.Left.Left
				} else if current.Right != nil {
					current.Value = current.Right.Value
					current.Left = current.Right.Left
					current.Right = current.Right.Right
				} else {

				}
			} else if parent.Left == current {
				if current.Left != nil {
					parent.Left = current.Left
				} else {
					parent.Left = current.Right
				}
			} else if parent.Right == current {
				if current.Left != nil {
					parent.Right = current.Left
				} else {
					parent.Right = current.Right
				}
			}
			break
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}
