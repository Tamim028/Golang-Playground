package main

import "fmt"

func structQueueTest()  {
	qint := queueInt{}

	qint.enque(5)
	qint.enque(6)
	qint.enque(7)

	fmt.Println(qint.list)

	qint.deque()

	fmt.Println(qint.list)

	qint.deque()

	fmt.Println(qint)

	qint.deque()
	qint.deque()
	qint.enque(798)
	qint.enque(78)

	fmt.Println(qint.list)


	qstr := queueString{}

	qstr.enque("Hello")
	qstr.enque("there")
	fmt.Println(qstr)
}

func sliceQueueTest()  {
	q := []int{}

	q = enqueue(q, 5)
	q = enqueue(q, 9)
	q = enqueue(q, 4)

	fmt.Println(q)

	q = dequeue(q)
	q = dequeue(q)

	fmt.Println(q)
}








