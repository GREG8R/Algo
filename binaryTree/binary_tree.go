package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Node struct{
	Value int
	LeftNode *Node
	RightNode *Node
}

func GenerateRandomTree(n int) *Node{
	arr := make([]int, n)
	for i := 0; i < n; i++{
		arr[i] = rand.Intn(n)
	}

	return BuildTree(arr)
}

func GenerateRandomSortTree(n int) *Node{
	arr := make([]int, n)
	for i := 0; i < n; i++{
		arr[i] = rand.Intn(n)
	}
	sort.Ints(arr)

	return BuildTree(arr)
}

func BuildTree(arr []int) *Node{
	root := InitTree(arr[0])

	for i := 1; i < len(arr); i++{
		root.Insert(arr[i])
	}

	return root
}

func InitTree(x int) *Node {
	return &Node{
		Value:     x,
		LeftNode:  nil,
		RightNode: nil,
	}
}

func (t *Node) Insert(x int, ){
	if t.Value > x {
		if t.LeftNode != nil {
			t.LeftNode.Insert(x)
		} else {
			t.LeftNode = &Node{
				Value: x,
				LeftNode:  nil,
				RightNode: nil,
			}
		}
	} else {
		if t.RightNode != nil {
			t.RightNode.Insert(x)
		} else {
			t.RightNode = &Node{
				Value: x,
				LeftNode:  nil,
				RightNode: nil,
			}
		}
	}
}


func Remove(x int, root *Node){
	nodeForDelete := root
	var parentNodeForDelete *Node = nil

	for nodeForDelete != nil {
		if nodeForDelete.Value == x{
			break
		} else {
			parentNodeForDelete = nodeForDelete
			if nodeForDelete.Value > x {
				nodeForDelete = nodeForDelete.LeftNode
			} else {
				nodeForDelete = nodeForDelete.RightNode
			}
		}
	}

	if parentNodeForDelete == nil {
		if nodeForDelete != nil && nodeForDelete.LeftNode != nil {
			root = nodeForDelete.LeftNode
		}
		if nodeForDelete != nil && nodeForDelete.RightNode != nil {
			root.TreeInsert(nodeForDelete.RightNode)
		}
	} else {
		if parentNodeForDelete.RightNode == nodeForDelete {
			parentNodeForDelete.RightNode = nil
		} else {
			parentNodeForDelete.LeftNode = nil
		}

		if nodeForDelete != nil && nodeForDelete.LeftNode != nil {
			parentNodeForDelete.TreeInsert(nodeForDelete.LeftNode)
		}

		if nodeForDelete != nil && nodeForDelete.RightNode != nil {
			parentNodeForDelete.TreeInsert(nodeForDelete.RightNode)
		}
	}

}

func (t *Node) TreeInsert(tree *Node){
	if tree.LeftNode != nil{
		t.TreeInsert(tree.LeftNode)
	}
	t.Insert(tree.Value)
	if tree.RightNode != nil {
		t.TreeInsert(tree.RightNode)
	}
}

func (t *Node) Search(x int) bool {
	if t == nil {
		return false
	}

	if t.Value == x {
		return true
	}

	if t.Value > x{
		return t.LeftNode.Search(x)
	}

	if t.Value < x{
		return t.RightNode.Search(x)
	}

	return false
}

func (t *Node) TreePrint(){
	if t.LeftNode != nil{
		t.LeftNode.TreePrint()
	}
	fmt.Println(t.Value)
	if t.RightNode != nil {
		t.RightNode.TreePrint()
	}
}

func (t *Node) TreeValues(array *[]int){
	if t.LeftNode != nil{
		t.LeftNode.TreeValues(array)
	}
	*array = append(*array, t.Value)
	if t.RightNode != nil {
		t.RightNode.TreeValues(array)
	}
}

func RemoveWithoutReqursion(val int, root *Node){
	nodeForDelete := root
	var parentNodeForDelete *Node = nil

	for nodeForDelete != nil {
		if val == nodeForDelete.Value{
			break
		} else {
			parentNodeForDelete = nodeForDelete
			if nodeForDelete.Value > val {
				nodeForDelete = nodeForDelete.LeftNode
			} else {
				nodeForDelete = nodeForDelete.RightNode
			}
		}
	}

	if nodeForDelete == nil {
		return
	}

	if nodeForDelete.RightNode == nil {
		if parentNodeForDelete == nil {
			root = nodeForDelete.LeftNode
		} else {
			if nodeForDelete == parentNodeForDelete.LeftNode{
				parentNodeForDelete.LeftNode = nodeForDelete.LeftNode
			} else {
				parentNodeForDelete.RightNode = nodeForDelete.LeftNode
			}
		}
	} else {
		leftMost := nodeForDelete.RightNode
		var parentNodeForCurrent *Node = nil
		for leftMost.LeftNode != nil {
			parentNodeForCurrent = leftMost
			leftMost = leftMost.LeftNode
		}

		if parentNodeForCurrent != nil {
			parentNodeForCurrent.LeftNode = leftMost.RightNode
		} else {
			nodeForDelete.RightNode = leftMost.RightNode
		}

		nodeForDelete.Value = leftMost.Value
	}

}
