package serviceB

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "project/queue"
)

var (
    mq          *queue.MessageQueue
    messageChan chan string
)

func handlerB(w http.ResponseWriter, r *http.Request) {
    select {
    case msg := <-messageChan:
        log.Printf("Service B: Processed message: %s\n", msg)
        fmt.Fprintf(w, "Processed message: %s", msg)
    default:
        time.Sleep(100 * time.Millisecond)
        select {
        case msg := <-messageChan:
            log.Printf("Service B: Processed message after retry: %s\n", msg)
            fmt.Fprintf(w, "Processed message: %s", msg)
        default:
            log.Println("Service B: No messages in queue")
            http.Error(w, "No messages in queue", http.StatusNoContent)
        }
    }
}

func StartServiceB(queue *queue.MessageQueue) {
    mq = queue
    messageChan = make(chan string, 10)

    go func() {
        for message := range mq.Subscribe() {
            log.Printf("Service B: Received message from queue: %s\n", message.Content)
            messageChan <- message.Content
        }
    }()

    http.HandleFunc("/service-b", handlerB)
    log.Println("Service B listening on port 8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
