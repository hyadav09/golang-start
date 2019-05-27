package main

import "fmt"

func main() {
	var sarray = []int{453, 22, 54, 12, 294, 67, -1, 46, 90, 0}
	var i, j, k int
	var n = len(sarray)
	var temp int

	for j = 1; j < n; j++ {
		for i = 0; i < j; i++ {
			if sarray[j] < sarray[i] {
				temp = sarray[j]
				for k = j; k > i; k-- {
					sarray[k] = sarray[k-1]
				}
				sarray[i] = temp
			}

		}

	}
	for k = 0; k < n; k++ {
		fmt.Printf("arr[%d] = %d, ", k, sarray[k])
	}
}
