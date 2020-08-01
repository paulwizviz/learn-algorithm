package main

import (
	"fmt"
)

func Example_ops_1() {
	var queue []int

	fmt.Printf("Initial: %v\n", queue)
	queue = enqueue(queue, 1)
	fmt.Printf("Add 1: %v\n", queue)
	queue = enqueue(queue, 2)
	fmt.Printf("Add 2: %v\n", queue)
	queue = enqueue(queue, 3)
	fmt.Printf("Add 3: %v\n", queue)
	queue = enqueue(queue, 4)
	fmt.Printf("Add 4: %v\n", queue)

	queue = dequeue(queue)
	fmt.Printf("First dequeue %v\n", queue)
	queue = dequeue(queue)
	fmt.Printf("Second dequeue %v\n", queue)

	// output:
	// Initial: []
	// Add 1: [1]
	// Add 2: [1 2]
	// Add 3: [1 2 3]
	// Add 4: [1 2 3 4]
	// First dequeue [2 3 4]
	// Second dequeue [3 4]
}

func Example_ops_2() {

	initElements := []int{1, 2, 3, 4, 5}
	var queue []int

	queue = initialise(queue, initElements)

	for i := 6; i < 21; i++ {
		queue = enqueue(queue, i)
	}
	fmt.Printf("Enqueue: %v\n", queue)

	for i := 1; i < 6; i++ {
		queue = dequeue(queue)
	}

	fmt.Printf("Dequeue: %v", queue)

	// output:
	// 	Enqueue: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	// Dequeue: [6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
}
