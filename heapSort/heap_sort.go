package main

func HeapSort(array []int) {
	swap := func(x, y int) {
		tmp := array[x]
		array[x] = array[y]
		array[y] = tmp
	}

	down := func(size, root int) {}
	down = func(size, root int) {
		l := root*2 + 1
		r := l + 1
		x := root
		if l < size && array[l] > array[x] {
			x = l
		}
		if r < size && array[r] > array[x] {
			x = r
		}
		if x == root {
			return
		}
		swap(root, x)
		down(size, x)
	}

	for node := len(array)/2 - 1; node >= 0; node-- {
		down(len(array), node)
	}

	down(len(array), 0)
	for i := len(array) - 1; i >= 0; i-- {
		swap(i, 0)
		down(i, 0)
	}
}
