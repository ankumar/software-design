package main

import (
    "fmt"
    
    "context"
    "log"
    
    "github.com/google/uuid"
    "go.temporal.io/sdk/client"
)

func main() {
    //fmt.Println("Hello, World!")
    
    clt, err := client.NewClient(client.options{})
    if err != nil {
        log.Fatalln("Temporal NewClient failed", err)
    }
    defer clt.Close()
    
}
