package main

import (
	"fmt"
	"net/http"
)

func main() {
	var URL string = "https://65c49fc3dae2304e92e2f2eb.mockapi.io/api/todoList"
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Some error!")
	}
	fmt.Println((resp))
}
