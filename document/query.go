package document

import (
	"sort"
	"time"
	"example.com/go67/operation"
)

func (s *Store) ListAfter(at time.Time, id string, limit int) []operation.PatchOperation {
	s.mu.RLock(); defer s.mu.RUnlock()
	out := make([]operation.PatchOperation, 0, len(s.items))
	for _, e := range s.items {
		if e.UpdatedAt.After(at) || (e.UpdatedAt.Equal(at) && e.ID > id) { out = append(out, e.Clone()) }
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].UpdatedAt.Equal(out[j].UpdatedAt) { return out[i].ID < out[j].ID }
		return out[i].UpdatedAt.Before(out[j].UpdatedAt)
	})
	if limit >= 0 && len(out) > limit { out = out[:limit] }
	return out
}

func (s *Store) DeleteBefore(cutoff time.Time) int {
	s.mu.Lock(); defer s.mu.Unlock()
	removed := 0
	for id, e := range s.items {
		if e.UpdatedAt.Before(cutoff) { delete(s.items,id); removed++ }
	}
	return removed
}
