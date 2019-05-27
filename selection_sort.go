package main

import "fmt"

func main() {
	var sarray = []int{453, 22, 54, 12, 294, 67, -1, 46, 90, 0}
	var i, j, k int
	var n = len(sarray)

	var smallest, index int

	for i = 0; i < n-1; i++ {
		smallest = sarray[i]
		index = i
		for j = i + 1; j < n; j++ {
			if smallest > sarray[j] {
				smallest = sarray[j]
				index = j
			}
		}
		if index != i {
			sarray[index] = sarray[i]
			sarray[i] = smallest
		}

	}
	for k = 0; k < n; k++ {
		fmt.Printf("arr[%d] = %d, ", k, sarray[k])
	}
}
