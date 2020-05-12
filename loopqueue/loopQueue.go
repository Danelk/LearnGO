package loopqueue

import "errors"

type LoopQueue struct {
	front 	int
	tail	int
	len 	int
	cap 	int
	queue 	[]interface{}
}

func NewLoopQueue() *LoopQueue {
	loop := &LoopQueue{
		front: 0,
		tail: 0,
		len: 0,
		cap: 3,
		queue: make([]interface{}, 0, 3),
	}
	for i:= 0; i<3; i++ {
		loop.queue = append(loop.queue, "")
	}
	return loop
}
func (q *LoopQueue) Len() int {
	return q.len
}
func (q *LoopQueue) Cap() int {
	return q.cap
}
func (q *LoopQueue) IsEmpty() bool {
	return q.len == 0
}
func (q *LoopQueue) IsFull() bool {
	return q.len+1 == q.cap
}

func (q *LoopQueue) GetFront() (interface{}, error) {
	if q.len == 0 {
		return nil, errors.New("Queue is Empty")
	}
	return q.queue[q.front], nil
}

