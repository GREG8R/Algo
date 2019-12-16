package main

// shell sort with Sedgewick gaps
func ShellSort(array []int) {
	gaps := []int{4197377, 1050113, 262913, 65921, 16577, 4193, 1073, 281, 77, 23, 8, 1}
	n := len(array)
	for _, gap := range gaps {
		for i := 0; i+gap < n; i++ {
			j := i + gap
			tmp := array[j]
			for ; j-gap >= 0 && array[j-gap] > tmp; j -= gap {
				array[j] = array[j-gap]
			}
			array[j] = tmp
		}
	}
}
