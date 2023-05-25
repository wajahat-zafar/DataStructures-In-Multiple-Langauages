package main
import "fmt"

type Node struct {
   value     int
   left_val  *Node  //create left_val and right_val elements 
   right_val *Node
}

func main() {
   root := &Node{value: 40} 
   root.left_val = &Node{value: 20}
   root.right_val = &Node{value: 60}
   root.left_val.left_val = &Node{value: 10}
   root.left_val.right_val = &Node{value: 30}
   root.right_val.left_val = &Node{value: 50}
   root.right_val.right_val = &Node{value: 70}
   
   fmt.Println("Nodes are as follows:")
   inOrder(root)

   deleteNode(root, 50)
   fmt.Println("Nodes after deletion:")
   inOrder(root)
   sortedValues := sort(root)
   fmt.Println("Nodes in sorted order:")
   for _, value := range sortedValues {
	   fmt.Println(value)
   }

   searchValue := 40
	found := search(root, searchValue)
	if found {
		fmt.Println(searchValue, "is found in the BST.")
	} else {
		fmt.Println(searchValue, "is not found in the BST.")
	}
}

func inOrder(node *Node) {
   if node == nil {
      return
   }
   inOrder(node.left_val)
   fmt.Println(node.value)
   inOrder(node.right_val)
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return root
	}

	if value < root.value {
		root.left_val = deleteNode(root.left_val, value)
	} else if value > root.value {
		root.right_val = deleteNode(root.right_val, value)
	} else {
		
		if root.left_val == nil {
			return root.right_val
		}
		
		if root.right_val == nil {
			return root.left_val
		}
		
		minValue := minValue(root.right_val)
		root.value = minValue
		root.right_val = deleteNode(root.right_val, minValue)
	}

	return root
}

func minValue(node *Node) int {
	min := node.value
	for node.left_val != nil {
		min = node.left_val.value
		node = node.left_val
	}
	return min
}

func sort(node *Node) []int {
	var sorted []int
	inOrderSort(node, &sorted)
	return sorted
}

func inOrderSort(node *Node, sorted *[]int) {
	if node == nil {
		return
	}
	inOrderSort(node.left_val, sorted)
	*sorted = append(*sorted, node.value)
	inOrderSort(node.right_val, sorted)
}

func search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	} else if value < node.value {
		return search(node.left_val, value)
	} else {
		return search(node.right_val, value)
	}
}