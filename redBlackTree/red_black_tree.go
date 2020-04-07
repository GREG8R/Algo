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
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}

	return Build(arr)
}

func GenerateSortRB(n int) *RedBlackTree {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	sort.Ints(arr)

	return Build(arr)
}

func Build(arr []int) *RedBlackTree {
	tree := &RedBlackTree{
		Root: nil,
	}

	for i := 0; i < len(arr); i++ {
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
	if node.Value > x {
		return node.Left.Search(x)
	}
	if node.Value < x {
		return node.Right.Search(x)
	}
	return false
}

func (node *RedBlackNode) Values(array *[]int) {
	if node.Left != nil {
		node.Left.Values(array)
	}
	*array = append(*array, node.Value)
	if node.Right != nil {
		node.Right.Values(array)
	}
}
