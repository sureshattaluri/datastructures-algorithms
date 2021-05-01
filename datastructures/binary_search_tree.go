package datastructures

import "fmt"

type BinarySearchTree struct {
	Root   *BstNode
	Length int
}

type BstNode struct {
	Value int
	Left  *BstNode
	Right *BstNode
}

func (bst *BinarySearchTree) Insert(in int) {

	v := &BstNode{
		Value: in,
		Left: nil,
		Right: nil,
	}

	if bst.Root == nil {
		bst.Root = v
	} else {
		currentNode := bst.Root
		for currentNode != nil {
			if in < currentNode.Value {
				//left
				if (currentNode.Left == nil) {
					currentNode.Left = v
				}
				currentNode = currentNode.Left
			} else {
				if (currentNode.Right == nil) {
					currentNode.Right = v
				}
				currentNode = currentNode.Right
			}
		}
	}
	bst.Length = bst.Length + 1
}

func (bst *BinarySearchTree) Lookup(in int) {
	currentNode := bst.Root
	iter := 0;
	for currentNode != nil {
		fmt.Printf("\n current node value: %v", currentNode.Value)
		if currentNode.Value == in {
			fmt.Printf("\n found value: %v after %v iterations", in, iter)
		}
		if in < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
		iter = iter + 1
	}
}

func TestBinarySearchTree () {
	var bst BinarySearchTree

	bst.Insert(9)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(4)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	bst.Lookup(1)
}
