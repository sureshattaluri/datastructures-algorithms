package main

import (
	"fmt"
	"github.com/sureshattaluri/datastructures-algorithms/linkedlist"
)

func main() {
	var ll linkedlist.LinkedList
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

	var llInsert linkedlist.LinkedList

	llInsert.Insert(10)
	llInsert.Insert(20)
	llInsert.Insert(5)
	llInsert.Insert(30)
	llInsert.Insert(15)

	currentNode := ll.Head;
	for currentNode != nil {
		fmt.Printf("\n %v", currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("\n ##################")
	currentNode = llInsert.Head;
	for currentNode != nil {
		fmt.Printf("\n %v", currentNode.Value)
		currentNode = currentNode.Next
	}

}
