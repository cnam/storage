package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestNewStorage test create new storage instance
func TestNewStorage(t *testing.T) {
	s := NewStorage()
	assert.IsType(t, &Storage{}, s)
}

// TestStorage_Get test get key from storage
func TestStorage_Get(t *testing.T) {
	s := NewStorage()
	s.data["key"] = "value"
	v, err := s.Get("key")
	assert.Nil(t, err)
	assert.Equal(t, "value", v)
}

// TestStorage_Set test set key to storage
func TestStorage_Set(t *testing.T) {
	s := NewStorage()
	n, err := s.Set("key", "value")
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, "value", s.data["key"])
}

// TestStorage_Del test del key from storage
func TestStorage_Del(t *testing.T) {
	s := NewStorage()
	s.data["key"] = "value"
	ok := s.Del("key")
	assert.True(t, ok)
	assert.Equal(t, s.data["key"], "")
}

// TestStorage_Keys test get all keys from storage
func TestStorage_Keys(t *testing.T) {
	s := NewStorage()
	s.data["key"] = "value"
	s.data["key1"] = "value1"
	k, err := s.Keys()
	assert.Nil(t, err)
	assert.Equal(t, []string{"key", "key1"}, k)
}
