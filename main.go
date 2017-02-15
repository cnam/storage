package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(fmt.Sprintf("Error start: %s", err))
	}
	s := NewStorage()
	ListenAndServe(Handler{ln, s})
}
