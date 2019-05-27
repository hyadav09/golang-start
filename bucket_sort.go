package main

import "fmt"

type node struct {
	value float64
	next  *node
}

func main() {
	var sarray = []float64{0.23, 0.98, 0.19, 0.45, 0.67, 0.11, 0.34, 0.33, 0.67, 0.44}
	var n = len(sarray)
	//var aptr [10]*node
	aptr := make([]*node, n)
	var i int

	for i = 0; i < n; i++ {
		index := int(sarray[i] * 10)
		aptr[index] = addNodeSorted(&node{sarray[i], nil}, aptr[index])
	}

	/*for i = 0; i < n; i++ {
		printlist(aptr[i])
	}*/

	index := 0
	var nd *node
	for i = 0; i < n; i++ {
		if aptr[i] != nil {
			for nd = aptr[i]; nd != nil; nd = nd.next {
				sarray[index] = nd.value
				index++
			}
		}
	}

	for i = 0; i < n; i++ {
		fmt.Println(sarray[i])
	}
}

func printlist(nodeList *node) {
	for nd := nodeList; nd != nil; nd = nd.next {
		fmt.Println(nd)
	}
	fmt.Println()
}

func addNodeSorted(newNode, nodeList *node) *node {
	if newNode == nil {
		return nodeList
	} else if nodeList == nil {
		nodeList = newNode
		return nodeList
	}

	for p := nodeList; p != nil; p = p.next {
		if newNode.value < p.value {
			newNode.next = p
			return newNode
		} else if p.next == nil {
			p.next = newNode
			return nodeList
		} else if newNode.value >= p.value && newNode.value < p.next.value {
			newNode.next = p.next
			p.next = newNode
			return nodeList
		}
	}

	return nodeList
}
