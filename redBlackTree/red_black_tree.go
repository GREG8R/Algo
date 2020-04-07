package main

import (
	"math/rand"
	"sort"
)

const Black int = 0
const Red int = 1

// RedBlackTree
type RedBlackTree struct {
	Root *RedBlackNode
}

type RedBlackNode struct {
	Left   *RedBlackNode
	Right  *RedBlackNode
	Parent *RedBlackNode

	Color int
	Value int
}

func GenerateRandomRB(n int) *RedBlackTree {
	arr := make([]int, n)
	for i := 0; i < n; i++{
		arr[i] = rand.Intn(n)
	}

	return Build(arr)
}

func GenerateSortRB(n int) *RedBlackTree {
	arr := make([]int, n)
	for i := 0; i < n; i++{
		arr[i] = rand.Intn(n)
	}
	sort.Ints(arr)

	return Build(arr)
}

func Build(arr []int) *RedBlackTree{
	tree := &RedBlackTree{
		Root: nil,
	}

	for i := 0; i < len(arr); i++{
		tree.Insert(arr[i])
	}

	return tree
}

func (t *RedBlackTree) LeftRotate(x *RedBlackNode) {
	y := x.Right
	x.Right = y.Left

	if y.Left != nil {
		y.Left.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == nil {
		t.Root = y
	} else if x.Parent.Left == x {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func (t *RedBlackTree) RightRotate(y *RedBlackNode) {
	x := y.Left
	y.Left = x.Right

	if x.Right != nil {
		x.Right.Parent = y
	}
	x.Parent = y.Parent

	if y.Parent == nil {
		t.Root = x
	} else if y.Parent.Right == y {
		y.Parent.Right = x
	} else {
		y.Parent.Left = x
	}

	x.Right = y
	y.Parent = x
}

func (t *RedBlackTree) Insert(key int) {
	x := &RedBlackNode{
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Color:  Black,
		Value:  key,
	}

	var parent *RedBlackNode = nil
	current := t.Root
	for current != nil {
		parent = current
		if x.Value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	x.Parent = parent
	x.Color = Red

	if parent == nil {
		t.Root = x
	} else if key < parent.Value {
		parent.Left = x
	} else {
		parent.Right = x
	}

	t.InsertFixup(x)
}


func (t *RedBlackTree) InsertFixup(x *RedBlackNode) {
	var y *RedBlackNode = nil

	for x != t.Root && x.Parent.Color == Red {
		if x.Parent == x.Parent.Parent.Left {
			y = x.Parent.Parent.Right
			if y != nil && y.Color == Red {
				x.Parent.Color = Black
				y.Color = Black
				x.Parent.Parent.Color = Red
				x = x.Parent.Parent
			} else if x == x.Parent.Right {
				x = x.Parent
				t.LeftRotate(x)
			} else {
				x.Parent.Color = Black
				x.Parent.Parent.Color = Red
				t.RightRotate(x.Parent.Parent)
			}
		} else {
			y = x.Parent.Parent.Left
			if y != nil && y.Color == Red {
				x.Parent.Color = Black
				y.Color = Black
				x.Parent.Parent.Color = Red
				x = x.Parent.Parent
			} else if x == x.Parent.Left {
				x = x.Parent
				t.RightRotate(x)
			} else {
				x.Parent.Color = Black
				x.Parent.Parent.Color = Red
				t.LeftRotate(x.Parent.Parent)
			}
		}
	}

	t.Root.Color = Black
}


func (node *RedBlackNode) Search(x int) bool {
	if node == nil {
		return false
	}
	if node.Value == x {
		return true
	}
	if node.Value > x{
		return node.Left.Search(x)
	}
	if node.Value < x{
		return node.Right.Search(x)
	}
	return false
}

func (node *RedBlackNode) Values(array *[]int){
	if node.Left != nil{
		node.Left.Values(array)
	}
	*array = append(*array, node.Value)
	if node.Right != nil {
		node.Right.Values(array)
	}
}

//func Print(node *RedBlackNode){
//	fmt.Printf("   %d %d", node.Value, node.Color)
//}
//
//func TreePrint(node *RedBlackNode){
//	if node.Left != nil {
//		Print(node.Left)
//	}
//	if node.Right != nil {
//		Print(node.Right)
//	}
//
//	if node.Left != nil {
//		TreePrint(node.Left)
//	}
//	if node.Right != nil {
//		TreePrint(node.Right)
//	}
//}

//// @param: node, a RedBlackNode
//// @param: node, the node with the smallest key rooted at node
//public RedBlackNode<T> treeMinimum(RedBlackNode<T> node){
//
//// while there is a smaller key, keep going left
//while (!isNil(node.left))
//node = node.left;
//return node;
//}// end treeMinimum(RedBlackNode node)

//func (t *RedBlackTree) RightRotateFixup (y *RedBlackNode){
//	if y.Right == nil && y.Left.Right == nil{
//		y.NumRight = 0
//		y.NumLeft = 0
//		y.Left.NumRight = 1
//	}	else if y.Right == nil && y.Left.Right != nil{
//		y.NumRight = 0
//		y.NumLeft = 1 + y.Left.Right.NumRight + y.Left.Right.NumLeft
//		y.Left.NumRight = 2 + y.Left.Right.NumRight + y.Left.Right.NumLeft
//	} else if y.Right != nil && y.Left.Right == nil{
//		y.NumLeft = 0
//		y.Left.NumRight = 2 + y.Right.NumRight + y.Right.NumLeft
//
//	}	else{
//		y.NumLeft = 1 + y.Left.Right.NumRight + y.Left.Right.NumLeft
//		y.Left.NumRight = 3 + y.Right.NumRight + y.Right.NumLeft + y.Left.Right.NumRight + y.Left.Right.NumLeft
//	}
//}

//func (t *RedBlackTree) LeftRotateFixup(x *RedBlackNode){
//	if x.Left == nil && x.Right.Left == nil{
//		x.NumLeft = 0
//		x.NumRight = 0
//		x.Right.NumLeft = 1
//	} else if x.Left == nil && x.Right.Left != nil{
//		x.NumLeft = 0
//		x.NumRight = 1 + x.Right.Left.NumLeft + x.Right.Left.NumRight
//		x.Right.NumLeft = 2 + x.Right.Left.NumLeft + x.Right.Left.NumRight
//	} else if x.Left != nil && x.Right.Left == nil{
//		x.NumRight = 0
//		x.Right.NumLeft = 2 + x.Left.NumLeft + x.Left.NumRight
//	} else{
//		x.NumRight = 1 + x.Right.Left.NumLeft + x.Right.Left.NumRight
//		x.Right.NumLeft = 3 + x.Left.NumLeft + x.Left.NumRight + x.Right.Left.NumLeft + x.Right.Left.NumRight
//	}
//}
