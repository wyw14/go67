package document

import (
	"sync"
	"example.com/go67/operation"
)

type Store struct {
	mu sync.RWMutex
	items map[string]operation.PatchOperation
}

func New() *Store { return &Store{items: make(map[string]operation.PatchOperation)} }

func (s *Store) Save(e operation.PatchOperation) {
	s.mu.Lock(); defer s.mu.Unlock()
	s.items[e.ID] = e.Clone()
}

func (s *Store) Get(id string) (operation.PatchOperation, bool) {
	s.mu.RLock(); defer s.mu.RUnlock()
	e, ok := s.items[id]
	return e.Clone(), ok
}
