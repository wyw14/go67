package document

import (
	"errors"
	"example.com/go67/operation"
)

var ErrNotFound = errors.New("patcher item not found")

func (s *Store) Update(id string, fn func(*operation.PatchOperation) error) error {
	s.mu.Lock(); defer s.mu.Unlock()
	current, ok := s.items[id]
	if !ok { return ErrNotFound }
	work := current.Clone()
	if err := fn(&work); err != nil { return err }
	s.items[id] = work.Clone()
	return nil
}
