package queue

import (
    "sync"
)

type Message struct {
    Content string
}

type MessageQueue struct {
    queue chan Message
    mu    sync.Mutex
}

func NewMessageQueue(size int) *MessageQueue {
    return &MessageQueue{
        queue: make(chan Message, size),
    }
}

func (mq *MessageQueue) Publish(message Message) {
    mq.mu.Lock()
    defer mq.mu.Unlock()
    mq.queue <- message
}

func (mq *MessageQueue) Subscribe() <-chan Message {
    return mq.queue
}
