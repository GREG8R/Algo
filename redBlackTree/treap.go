package main

import (
	"math/rand"
	"sort"
)

type Treap struct{
	x int
	y int
	Left *Treap
	Right *Treap
}

func GenerateRandom(n int) *Treap {
	arr := make([]int, n)
	for i := 0; i < n; i++{
		arr[i] = rand.Intn(n)
	}

	return BuildTreap(arr)
}

func GenerateSort(n int) *Treap {
	arr := make([]int, n)
	for i := 0; i < n; i++{
		arr[i] = rand.Intn(n)
	}
	sort.Ints(arr)

	return BuildTreap(arr)
}

func BuildTreap(arr []int) *Treap{
	root := InitTreap(arr[0], rand.Intn(100000))

	for i := 1; i < len(arr); i++{
		root = root.Add(arr[i])
	}

	return root
}

func InitTreap(x int, y int) *Treap {
	return &Treap{ x, y, nil, nil }
}

func Merge(l, r *Treap) *Treap {
	if l == nil { return r }
	if r == nil { return l }

	if l.y > r.y {
		newR := Merge(l.Right, r)
		return &Treap{l.x, l.y, l.Left, newR}
	} else {
		newL := Merge(l, r.Left)
		return &Treap{r.x, r.y, newL, r.Right}
	}
}

func (t *Treap) Split(x int) (l, r *Treap){
	var newTree *Treap = nil
	if t.x <= x {
		if t.Right == nil {
			r = nil
		} else {
			newTree, r = t.Right.Split(x)
		}
		l = &Treap{ t.x, t.y, t.Left, newTree}
	} else {
		if t.Left == nil{
			l = nil
		} else {
			l, newTree = t.Left.Split(x)
		}
		r = &Treap{ t.x, t.y, newTree, t.Right}
	}
	return l, r
}

func (t *Treap) Add(x int) *Treap {
	l, r := t.Split(x)
	m := &Treap{x, rand.Intn(100000), nil, nil}
	return Merge(Merge(l, m), r)
}

func (t *Treap) Remove(x int) *Treap {
	l, r := t.Split(x - 1)
	_, r = r.Split(x)
	return Merge(l, r)
}

func (t *Treap) Search(x int) bool {
	if t == nil {
		return false
	}
	if t.x == x {
		return true
	}
	if t.x > x{
		return t.Left.Search(x)
	}
	if t.x < x{
		return t.Right.Search(x)
	}
	return false
}

func (t *Treap) TreapValues(array *[]int){
	if t.Left != nil{
		t.Left.TreapValues(array)
	}
	*array = append(*array, t.x)
	if t.Right != nil {
		t.Right.TreapValues(array)
	}
}