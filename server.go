package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Handler struct {
	net.Listener
	*Storage
}

// ListenAndServe
func ListenAndServe(handler Handler) {
	for {
		conn, err := handler.Accept()

		if err != nil {
			log.Printf("Error accept connection %s", conn)
		}

		go handler.accept(conn, handler.Storage)
	}
}

// accept new connection for work
func (h Handler) accept(conn net.Conn, s *Storage) {
	var cmd, key, value string

	for {
		fmt.Fscanf(conn, "%s %s %s", &cmd, &key, &value)

		switch strings.ToLower(cmd) {
		case "get":
			r, _ := s.Get(key)
			fmt.Fprintf(conn, "%s\n", r)
		case "set":
			s.Set(key, value)
		case "keys":
			k, _ := s.Keys()
			for _, v := range k {
				fmt.Fprintf(conn, "%s\n", v)
			}
		case "del":
			s.Del(key)
		case "exit":
			conn.Close()
			return
		}
	}
}
