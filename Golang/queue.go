package main

import "fmt"

func enqueue(queue []int, value int) []int {
	queue = append(queue, value)
	return queue
}

func dequeue(queue []int) (int, []int) {
	if len(queue) == 0 {
		fmt.Println("The queue is empty")
		return 0, queue
	}

	front := queue[0]
	queue = queue[1:]
	return front, queue
}

func sortQueue(queue []int) []int {
	for i := 0; i < len(queue)-1; i++ {
		for j := 0; j < len(queue)-i-1; j++ {
			if queue[j] > queue[j+1] {
				queue[j], queue[j+1] = queue[j+1], queue[j]
			}
		}
	}
	return queue
}

func main() {
	myQueue := []int{}
	myQueue = enqueue(myQueue, 2)
	myQueue = enqueue(myQueue, 3)
	myQueue = enqueue(myQueue, 4)
	myQueue = enqueue(myQueue, 5)
	myQueue = enqueue(myQueue, 3)
	myQueue = enqueue(myQueue, 8)
	myQueue = enqueue(myQueue, 14)
	myQueue = enqueue(myQueue, 9)

	fmt.Println(myQueue)

	front, myQueue2 := dequeue(myQueue)
	fmt.Println("Dequeued element:", front)
	fmt.Println("Remaining queue:", myQueue2)

	sortedQueue := sortQueue(myQueue)
	fmt.Println("The sorted queue would be:", sortedQueue)

	var ask int
	fmt.Println("Enter the number you want to search: ")
	fmt.Scanln(&ask)

	found := false
	for i := 0; i < len(myQueue); i++ {
		if myQueue[i] == ask {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Your searched number is present in the queue.")
	} else {
		fmt.Println("Your searched number is not present in the queue.")
	}
}
