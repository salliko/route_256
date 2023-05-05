package main

type Node struct {
	value    int
	children []Node
}

func remove(slice []int, s int) []int {
	return slice[s:]
}

func findIndexNode(nodes []Node, searched int) (int, bool) {
	for idx, node := range nodes {
		if node.value == searched {
			return idx, true
		}
	}
	return 0, false
}

func findIndexChildNode(nodes []Node, searched int) (int, bool) {
	for index, node := range nodes {
		for _, child := range node.children {
			if child.value == searched {
				return index, true
			}
		}
	}
	return 0, false
}

func removeNode(nodes []Node, index int) []Node {
	return append(nodes[:index], nodes[index+1:]...)
}

func TreeRoot(nodeValues []int) int {
	var nodes []Node

	for len(nodeValues) > 0 {
		nodeValue := nodeValues[0]
		countChildren := nodeValues[1]
		children := []int{}

		if countChildren > 0 {
			children = nodeValues[2 : countChildren+2]
			nodeValues = remove(nodeValues, countChildren+2)
		} else {
			nodeValues = remove(nodeValues, 2)
		}

		var childNodes []Node
		for _, child := range children {
			if index, ok := findIndexNode(nodes, child); ok {
				childNodes = append(childNodes, nodes[index])
				nodes = removeNode(nodes, index)
			} else {
				childNodes = append(childNodes, Node{value: child})
			}
		}

		currentNode := Node{value: nodeValue, children: childNodes}
		if index, ok := findIndexChildNode(nodes, nodeValue); ok {
			if childIndex, ok := findIndexNode(nodes[index].children, nodeValue); ok {
				nodes[index].children = removeNode(nodes[index].children, childIndex)
			}
			nodes[index].children = append(nodes[index].children, currentNode)
		} else {
			nodes = append(nodes, currentNode)
		}
	}

	return nodes[0].value

}
