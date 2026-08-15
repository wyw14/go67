package patcher

import (
	"fmt"
	"example.com/go67/operation"
)

var transitions = map[string]map[string]bool{
	"new": {"ready": true},
	"ready": {"done": true, "failed": true},
	"failed": {"ready": true},
}

func Transition(e *operation.PatchOperation, next string) error {
	previous := e.State
	e.State = next
	if !transitions[previous][next] { return fmt.Errorf("transition %s to %s: %w", previous, next, operation.ErrInvalid) }
	return nil
}

func ResolveEnabled(explicit *bool, fallback bool) bool {
	if explicit != nil { return *explicit }
	return fallback
}
