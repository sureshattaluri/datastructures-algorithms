package main

import (
	"fmt"
	as "github.com/sureshattaluri/datastructures-algorithms/algorithms"
	ds "github.com/sureshattaluri/datastructures-algorithms/datastructures"
)

func main() {
	//ds.TestLinkedList()
	ds.TestBinarySearchTree()

	fmt.Printf("Factorial of %v is %v", 5, as.Fact(5))
}
