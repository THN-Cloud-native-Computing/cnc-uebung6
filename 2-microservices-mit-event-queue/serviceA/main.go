package serviceA

import (
    "fmt"
    "log"
    "net/http"
    "project/queue"
)

var mq *queue.MessageQueue

func handlerA(w http.ResponseWriter, r *http.Request) {
    message := queue.Message{Content: "Message from Service A"}
    mq.Publish(message)
    log.Printf("Service A: Published message to the queue: %s\n", message.Content)
    fmt.Fprintf(w, "Published message to the queue: %s", message.Content)
}

func StartServiceA(queue *queue.MessageQueue) {
    mq = queue
    http.HandleFunc("/service-a", handlerA)
    log.Println("Service A listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
