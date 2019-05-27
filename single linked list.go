package main

import "fmt"

type numNode struct {
	value int
	next  *numNode
}

func main() {
	var numList *numNode

	n1 := &numNode{93, nil}
	numList = addNodeSorted(n1, nil)
	printList(numList)

	n2 := &numNode{34, nil}
	numList = addNodeSorted(n2, numList)
	printList(numList)

	n3 := &numNode{15, nil}
	numList = addNodeSorted(n3, numList)
	printList(numList)

	n4 := &numNode{56, nil}
	numList = addNodeSorted(n4, numList)
	printList(numList)

	n5 := &numNode{111, nil}
	numList = addNodeSorted(n5, numList)
	printList(numList)

	n6 := &numNode{55, nil}
	numList = addNodeSorted(n6, numList)
	printList(numList)
}

func printList(numList *numNode) {
	for p := numList; p != nil; p = p.next {
		fmt.Println(p)
	}
	fmt.Println()
}

func addNodeSorted(newNum, numList *numNode) *numNode {
	if newNum == nil {
		return numList
	} else if numList == nil {
		numList = newNum
		return numList
	}

	for p := numList; p != nil; p = p.next {
		if newNum.value < p.value {
			newNum.next = p
			return newNum
		} else if p.next == nil {
			p.next = newNum
			return numList
		} else if newNum.value >= p.value && newNum.value < p.next.value {
			newNum.next = p.next
			p.next = newNum
			return numList
		}
	}

	return numList
}
	