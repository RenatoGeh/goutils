// A simple FIFO implementation of type interface{}.
package queue

type Queue struct {
	data []interface{}
}

// Enqueue inserts element e at the end of the queue.
func (q *Queue) Enqueue(e interface{}) {
	q.data = append(q.data, e)
}

// Dequeue removes and returns the first element of the queue.
func (q *Queue) Dequeue() interface{} {
	n := len(q.data)
	e := q.data[0]
	q.data[0] = nil
	q.data = q.data[1:n]
	return e
}

// Peek returns the first element of the queue.
func (q *Queue) Peek() interface{} {
	return q.data[0]
}

// Get returns the i-th element of the queue.
func (q *Queue) Get(i int) interface{} {
	return q.data[i]
}

// Size returns the size of the queue.
func (q *Queue) Size() int { return len(q.data) }

// Empty returns whether the queue is empty or not.
func (q *Queue) Empty() bool { return len(q.data) == 0 }
