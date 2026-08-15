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
	if !transitions[e.State][next] { return fmt.Errorf("transition %s to %s: %w", e.State, next, operation.ErrInvalid) }
	e.State = next
	return nil
}

func ResolveEnabled(explicit *bool, fallback bool) bool {
	if explicit != nil { return *explicit }
	return fallback
}
