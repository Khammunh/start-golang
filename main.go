package main

import ("fmt"
"github.com/google/uuid"
)

func main() {
id := uuid.New()
fmt.Println("Hello Golang")
fmt.Println("UUID: %s", id)
}