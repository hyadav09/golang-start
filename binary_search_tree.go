package main

type tnode struct {
	value int
	lnode *tnode
	rnode *tnode
}

func main() {
	var sarray = []int{56, 12, 89, 1, 99, 32, 45, 78, 98}
	var n = len(sarray)
	var root *tnode

	root = addNode(&tnode{sarray[0], nil, nil}, nil)

	for i := 1; i < n; i++ {
		root = addNode(&tnode{sarray[i], nil, nil}, root)
	}

	printTree(root)
}

func printTree(root *tnode) {
	if root != nil {
		//println(root.value)
		printTree(root.lnode)
		//println(root.value)
		printTree(root.rnode)
		println(root.value)
	}
}

func addNode(newNode, root *tnode) *tnode {
	if newNode == nil {
		return root
	} else if root == nil {
		root = newNode
		return root
	}

	if newNode.value < root.value {
		if root.lnode != nil {
			root.lnode = addNode(newNode, root.lnode)
			return root
		} else {
			root.lnode = newNode
			return root
		}
	}

	if newNode.value > root.value {
		if root.rnode != nil {
			root.rnode = addNode(newNode, root.rnode)
			return root
		} else {
			root.rnode = newNode
			return root
		}
	}

	return root
}
