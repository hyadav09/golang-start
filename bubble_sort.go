package main

import "fmt"

func main() {
	var sarray = []int{4536705, 22, 54, 12, 294874, 67, -1, 46, 90, 0}
	var i, j, k, loop, tmp int
	var n = len(sarray)
	var cnt int
	for loop = 0; loop < n; loop++ {
		cnt = 0
		for i = 0; i < n-loop-1; i++ {
			j = i + 1
			if sarray[j] < sarray[i] {
				tmp = sarray[i]
				sarray[i] = sarray[j]
				sarray[j] = tmp
				cnt = cnt + 1
			}

		}
		for k = 0; k < n; k++ {
			fmt.Printf("arr[%d] = %d, ", k, sarray[k])
		}
		fmt.Printf("\n")
		if cnt == 0 {
			break
		}
	}
	for i = 0; i < n; i++ {
		fmt.Printf("sarray= %d, ", sarray[i])
	}
}
