package main

import "sync"

// Storage represents storage model
type Storage struct {
	sync.Mutex
	data map[string]string
}

// NewStorage create new storage object
func NewStorage() (*Storage) {
	return &Storage{
		data: make(map[string]string),
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
