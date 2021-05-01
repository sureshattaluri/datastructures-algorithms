package linkedlist

import "fmt"

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Node struct {
	Value    int32
	Next     *Node
	Previous *Node
}

func (ll *LinkedList) Insert(value int32) {

	v := &Node{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}

	if ll.Head == nil {
		ll.Head = v
	} else {
		currentNode := ll.Head
		var lastNode *Node
		fmt.Printf("\n current value: %v", v.Value)
		for currentNode != nil {
			fmt.Printf("\n current node value: %v", currentNode.Value)
			if currentNode.Value < value {
				changeNode := currentNode
				v.Next = changeNode
				v.Previous = changeNode.Previous
				changeNode.Previous = v

				if v.Previous == nil {
					ll.Head = v
				} else {
					v.Previous.Next = v
				}
				break
			}
			if currentNode.Next == nil {
				lastNode = currentNode
			}
			currentNode = currentNode.Next
		}
		if lastNode != nil {
			v.Previous = lastNode
			lastNode.Next = v
		}
	}
	ll.Length = ll.Length + 1
}

func (ll *LinkedList) Append(value int32) {

	v := &Node{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}

	if ll.Head == nil {
		ll.Head = v
		ll.Tail = v
	} else {
		ll.Tail.Next = v
		ll.Tail = v
	}
	ll.Length = ll.Length + 1
}

func (ll *LinkedList) Prepend(value int32) {

	v := &Node{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}

	if ll.Head == nil {
		ll.Head = v
		ll.Tail = v
	} else {
		v.Next = ll.Head
		ll.Head = v
	}
	ll.Length = ll.Length + 1
}
