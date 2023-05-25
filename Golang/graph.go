package main

import (
	"fmt"
	"sort"
)

type Node struct {
	value int
	edges []*Node
}

func (n *Node) addEdges(node *Node) {
	n.edges = append(n.edges, node)
}

func (n *Node) removeEdge(node *Node) {
	for i, edge := range n.edges {
		if edge == node {
			n.edges = append(n.edges[:i], n.edges[i+1:]...)
			break
		}
	}
}

func (n *Node) sortEdges() {
	sort.Slice(n.edges, func(i, j int) bool {
		return n.edges[i].value < n.edges[j].value
	})
}

func (n *Node) search(value int) *Node {
	visited := make(map[*Node]bool)
	queue := []*Node{n}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.value == value {
			return node
		}

		for _, edge := range node.edges {
			if !visited[edge] {
				visited[edge] = true
				queue = append(queue, edge)
			}
		}
	}

	return nil
}


func main() {
	n1 := &Node{value: 10}
	n2 := &Node{value: 20}
	n3 := &Node{value: 30}
	n4 := &Node{value: 40}
	n5 := &Node{value: 50}
	n6 := &Node{value: 60}

	n1.addEdges(n2)
	n1.addEdges(n3)
	n2.addEdges(n5)
	n2.addEdges(n6)
	n6.addEdges(n4)
	n4.addEdges(n3)
	n3.addEdges(n2)
	n5.addEdges(n5)
	n4.addEdges(n1)

	n1.sortEdges()
	n2.sortEdges()
	n3.sortEdges()
	n4.sortEdges()
	n5.sortEdges()

	fmt.Println("The Graph is represented as:")
	fmt.Printf("Node %d -> %v\n", n1.value, getNodeValues(n1.edges))
	fmt.Printf("Node %d -> %v\n", n2.value, getNodeValues(n2.edges))
	fmt.Printf("Node %d -> %v\n", n3.value, getNodeValues(n3.edges))
	fmt.Printf("Node %d -> %v\n", n4.value, getNodeValues(n4.edges))
	fmt.Printf("Node %d -> %v\n", n5.value, getNodeValues(n5.edges))

	n2.removeEdge(n5)
	n4.removeEdge(n1)

	fmt.Println("The Graph after removing an edge:")
	fmt.Printf("Node %d -> %v\n", n1.value, getNodeValues(n1.edges))
	fmt.Printf("Node %d -> %v\n", n2.value, getNodeValues(n2.edges))
	fmt.Printf("Node %d -> %v\n", n3.value, getNodeValues(n3.edges))
	fmt.Printf("Node %d -> %v\n", n4.value, getNodeValues(n4.edges))
	fmt.Printf("Node %d -> %v\n", n5.value, getNodeValues(n5.edges))

	fmt.Println("The Graph after sorting edges:")
	fmt.Printf("Node %d -> %v\n", n1.value, getNodeValues(n1.edges))
	fmt.Printf("Node %d -> %v\n", n2.value, getNodeValues(n2.edges))
	fmt.Printf("Node %d -> %v\n", n3.value, getNodeValues(n3.edges))
	fmt.Printf("Node %d -> %v\n", n4.value, getNodeValues(n4.edges))
	fmt.Printf("Node %d -> %v\n", n5.value, getNodeValues(n5.edges))

	
	searchValue := 40
	result := n1.search(searchValue)
	if result != nil {
		fmt.Printf("Node with value %d found.\n", searchValue)
	} else {
		fmt.Printf("Node with value %d not found.\n", searchValue)
	}
}

func getNodeValues(nodes []*Node) []int {
	var values []int
	for _, node := range nodes {
		values = append(values, node.value)
	}
	return values
}
