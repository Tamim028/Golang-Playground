package main

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	return  queue
}

func dequeue(queue []int) []int {
	return  queue[1:]
}

