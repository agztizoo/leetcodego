package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"

    "github.com/gin-gonic/gin"
)

func main() {
    http.HandleFunc("/", IndexHandler)
    err := http.ListenAndServe("127.0.0.1:8080", nil)
    if err !=nil {
        return
    }
    // engine := gin.Default()
    // engine.Any("/", WebRoot)
    // err = engine.Run(":8080")
    // if err !=nil {
    //     return
    // }
}

func WebRoot(context *gin.Context) {
    context.String(http.StatusOK, "hello")
}

func IndexHandler(w http.ResponseWriter, r * http.Request) {
    fmt.Fprintln(w, "hello")
}
