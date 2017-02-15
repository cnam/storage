package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type Server struct {
}

// Storage represents storage model
type Storage struct {
	sync.Mutex
	data map[string]string
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(fmt.Sprintf("Error start: %s", err))
	}
	b := &Storage{
		data: make(map[string]string),
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Printf("Error accept connection %s", conn)
		}

		go handleFunc(conn, b)
	}
}

func handleFunc(conn net.Conn, b *Storage) {
	var cmd, key, value string

	for {
		_, err := fmt.Fscanf(conn, "%s %s %s", &cmd, &key, &value)

		if err != nil {
			log.Printf("Error read data: %s", err)
		}

		switch strings.ToLower(cmd) {
		case "get":
			r, _ := b.Get(key)
			fmt.Fprintf(conn, "%s\n", r)
		case "set":
			b.Set(key, value)
		case "keys":
			k, _ := b.Keys()
			for _, v := range k {
				fmt.Fprintf(conn, "%s\n", v)
			}
		case "del":
			b.Del(key)
		case "exit":
			conn.Close()
			return
		}
	}
}

// Del deletes value by key from storage
func (s *Storage) Del(k string) (ok bool) {
	s.Lock()
	delete(s.data, k)
	s.Unlock()
	return true
}

// Keys return key list
func (s *Storage) Keys() ([]string, error) {
	s.Lock()
	var keys []string
	for k, _ := range s.data {
		keys = append(keys, k)
	}
	s.Unlock()

	return keys, nil
}

// Set add new value to storage
func (s *Storage) Set(k string, v string) (int, error) {
	s.Lock()
	s.data[k] = v
	s.Unlock()

	return 1, nil
}

// Get gets data from storage
func (s *Storage) Get(k string) (string, error) {
	s.Lock()
	d, _ := s.data[k]
	s.Unlock()
	return d, nil
}
