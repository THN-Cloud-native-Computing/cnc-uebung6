package main

import (
    "project/queue"
    "project/serviceA"
    "project/serviceB"
)

func main() {
    mq := queue.NewMessageQueue(10)

    go serviceA.StartServiceA(mq)
    serviceB.StartServiceB(mq)
}

