package queue

import "errors"

var (
	ErrQueueEmpty      = errors.New("queue is empty")
	ErrIndexOutOfBound = errors.New("index is out of queue bound")
)

// CharQueue is a queue that holds a collection of chars.
type CharQueue struct {
	queue []string
}

// NewCharFromString creates a new CharQueue from a string.
// The queue is intended to receive only alphanumeric characters.
// The elements of the queue will be individual characters contained in the string.
func NewCharFromString(s string) *CharQueue {
	queue := make([]string, len(s))
	for i := range s {
		queue[i] = string(s[i])
	}

	return &CharQueue{queue}
}

func (q *CharQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

// Peek returns the first element of the queue, this operation does not change the state of the queue.
func (q *CharQueue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", ErrQueueEmpty
	}

	return q.queue[0], nil
}

func (q *CharQueue) Get(index int) (string, error) {
	if index >= q.Size() {
		return "", ErrIndexOutOfBound
	}

	return q.queue[index], nil
}

func (q *CharQueue) Size() int {
	return len(q.queue)
}

// Enqueue adds new element to the end of queue, element is assumed to be a single character in form of string.
func (q *CharQueue) Enqueue(char string) {
	q.queue = append(q.queue, char)
}

// Dequeue removes and returns the first element of the queue.
func (q *CharQueue) Dequeue() (string, error) {
	if q.IsEmpty() {
		return "", ErrQueueEmpty
	}

	char := q.queue[0]
	q.queue = q.queue[1:]

	return char, nil
}
