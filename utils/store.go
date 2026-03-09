package utils

import (
	"log"
	"sync"
	"time"
)

type storeValue struct {
	value     string
	expiresAt time.Time
}

type Store interface {
	get(key string) (string, bool)
	put(key, val string)
	del(key string)
	list() map[string]string
	delExpired()
}

type InMemoryStore struct {
	store map[string]storeValue
	mu    *sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{store: make(map[string]storeValue), mu: &sync.RWMutex{}}
}

func (s *InMemoryStore) get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.store[key]
	return val.value, ok
}

func (s *InMemoryStore) put(key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	end := now.Add(5 * time.Minute)
	s.store[key] = storeValue{value: val, expiresAt: end}
}

func (s *InMemoryStore) del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}

func (s *InMemoryStore) list() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make(map[string]string)
	for k, v := range s.store {
		result[k] = v.value
	}
	return result
}

func (s *InMemoryStore) delExpired() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	amount := 0
	for key, value := range s.store {
		if now.Sub(value.expiresAt) > 0 {
			amount += 1
			delete(s.store, key)
		}
	}
	log.Printf("deleted %d expired keys\n", amount)
}
