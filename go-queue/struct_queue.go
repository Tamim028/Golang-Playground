package main

type queueInt struct {
	list []int
}

type queueString struct {
	list []string
}

func (q *queueInt) enque(value int){
	q.list = append(q.list,value)
}

func (q *queueInt) deque(){
	if len(q.list) == 0 {
		return
	}
	q.list = q.list[1:]
}

func (q *queueString) enque(value string){
	q.list = append(q.list,value)
}

func (q *queueString) deque(){
	if len(q.list) == 0 {
		return
	}
	q.list = q.list[1:]
}