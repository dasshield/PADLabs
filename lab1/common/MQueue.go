package common

import (
	"sync"
)

type MQueue struct {
	Messages []ServerMessage `json:"messages"`
	mutex    sync.Mutex
}

func (q *MQueue) Push(data ServerMessage) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.Messages = append(q.Messages, data)
}

func NewQueue() *MQueue {
	return &MQueue{Messages: make([]ServerMessage, 0)}
}

func (q *MQueue) Pop() ServerMessage {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	length := len(q.Messages)

	if length == 0 {
		return ServerMessage{Type:"nil"}
	}

	msg := q.Messages[length - 1]
	q.Messages = q.Messages[:length - 1]
	return msg
}
