package chans

import (
    "fmt"
    "time"
)

type Token struct {}

func newWorker(id int, cur, next chan Token)  {
    for {
        token := <-cur
        fmt.Println(id + 1)
        time.Sleep(time.Second)
        next <- token
    }
}

func RunWorker() {
    tokens := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
    for i := 0; i < 4; i++ {
        go newWorker(i, tokens[i], tokens[(i + 1)%len(tokens)])
    }
    tokens[0] <- Token{}
    select {}
}
