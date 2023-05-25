package main

import (
    "fmt"
)

type Node struct {
    data int
    next *Node
}

type List struct {
    head *Node
}

func (l *List) add(value int) {
    newNode := &Node{data: value}

    if l.head == nil {
        l.head = newNode
        return
    }

    curr := l.head
    for curr.next != nil {
        curr = curr.next
    }

    curr.next = newNode
}

func (l *List) remove(value int) {
    if l.head == nil {
        return
    }

    if l.head.data == value {
        l.head = l.head.next
        return
    }

    curr := l.head
    for curr.next != nil && curr.next.data != value {
        curr = curr.next
    }

    if curr.next != nil {
        curr.next = curr.next.next
    }
}

func (l *List) sort() {
    if l.head == nil || l.head.next == nil {
        return 
    }

    sorted := false
    for !sorted {
        sorted = true
        prev := l.head
        curr := prev.next

        for curr != nil {
            if prev.data > curr.data {
                
                prev.data, curr.data = curr.data, prev.data
                sorted = false
            }
            prev = curr
            curr = curr.next
        }
    }
}

func (l *List) search(value int) bool {
    curr := l.head
    for curr != nil {
        if curr.data == value {
            return true 
        }
        curr = curr.next
    }
    return false 
}



func main() {
    list := &List{}
    list.add(89)
    list.add(55)
    list.add(34)
    list.add(98)
	fmt.Println("The initial List:")
	printList(list)

    
	list.sort()
    fmt.Println("Sorted List:")
	printList(list)
	value := 55
    if list.search(value) {
        fmt.Printf("Value %d exists in the list.\n", value)
    } else {
        fmt.Printf("Value %d does not exist in the list.\n", value)
    }
    

    list.remove(34)
    fmt.Println("List after removing 2: ")
    printList(list)

    list.remove(98)
    fmt.Println("List after removing 4: ")
    printList(list)
}

func printList(l *List) {
    curr := l.head
    for curr != nil {
        fmt.Printf("%d ", curr.data)
        curr = curr.next
    }
    fmt.Println()

}