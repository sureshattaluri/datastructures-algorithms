package datastructures

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

func TestLinkedList() {
	var ll LinkedList
	ll.Append(10)
	ll.Append(20)
	ll.Append(5)
	ll.Append(30)
	ll.Append(15)

	ll.Prepend(10)
	ll.Prepend(20)
	ll.Prepend(5)
	ll.Prepend(30)
	ll.Prepend(15)

	var llInsert LinkedList

	llInsert.Insert(10)
	llInsert.Insert(20)
	llInsert.Insert(5)
	llInsert.Insert(30)
	llInsert.Insert(15)

	currentNode := ll.Head
	for currentNode != nil {
		fmt.Printf("\n %v", currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("\n ##################")
	currentNode = llInsert.Head
	for currentNode != nil {
		fmt.Printf("\n %v", currentNode.Value)
		currentNode = currentNode.Next
	}
}
