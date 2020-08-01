package main

func initialise(queue []int, elements []int) []int {
	queue = append(queue, elements...)
	return queue
}

// Pushing element to a queue of int
func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

// Remove queue from end
func dequeue(queue []int) []int {
	return queue[1:]
}
