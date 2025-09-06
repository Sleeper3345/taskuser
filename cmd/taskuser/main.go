package main

import (
    "log"
    "taskuser/internal/server"
)

func main() {
    e := server.New()
    log.Fatal(e.Start(":8080"))
}
