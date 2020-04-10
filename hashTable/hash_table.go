package main

type node struct {
	key       int
	value     interface{}
	isDeleted bool
}

type HashTable struct {
	nodes    []node
	capacity int
	length   int
}

func Init(n int) *HashTable {
	return &HashTable{
		nodes:    make([]node, n),
		capacity: n,
		length:   0,
	}
}

func (ht *HashTable) Set(key int, value interface{}) {
	if ht.length + 1 == ht.capacity{
		scale(ht)
	}
	i := 0
	for {
		if i >= ht.capacity {
			scale(ht)
		}
		position := ht.hashFunc(key, i)
		if ht.nodes[position].value == nil || ht.nodes[position].isDeleted {
			ht.nodes[position].key = key
			ht.nodes[position].value = value
			ht.length++
			return
		}

		if ht.nodes[position].key == key {
			ht.nodes[position].value = value
			return
		}
		i++
	}
}

func (ht *HashTable) Get(key int) interface{} {
	i := 0
	firstDeletedPosition := -10
	for {
		position := ht.hashFunc(key, i)
		if ht.nodes[position].key == key {
			value := ht.nodes[position].value
			if firstDeletedPosition >= 0 {
				ht.swap(firstDeletedPosition, position)
			}
			return value
		}
		if ht.nodes[position].value == nil {
			return nil
		}
		if ht.nodes[position].isDeleted {
			firstDeletedPosition = position
		}
		i++
	}
}

func (ht *HashTable) hashFunc(key, count int) int{
	return ((key % ht.capacity) + count * (1 + (key % (ht.capacity - 1)))) % ht.capacity
}

func (ht *HashTable) Delete(key int) {
	i := 0
	for {
		position := ht.hashFunc(key, i)
		if ht.nodes[position].key == key{
			ht.nodes[position].isDeleted = true
			return
		}
		i++
	}
}

func (ht *HashTable) swap(pos1, pos2 int){
	buf := ht.nodes[pos1]
	ht.nodes[pos1] = ht.nodes[pos2]
	ht.nodes[pos2] = buf
}

func scale(ht *HashTable) {
	newCap := ht.capacity * 2
	newHt := Init(newCap)
	for _, n := range ht.nodes{
		if !n.isDeleted && n.value != nil {
			newHt.Set(n.key, n.value)
		}
	}

	ht.nodes = newHt.nodes
	ht.capacity = newHt.capacity
	ht.length = newHt.length
}
