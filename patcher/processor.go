package patcher

import (
	"context"
	"example.com/go67/operation"
)

func Process(ctx context.Context, values []operation.PatchOperation, deliver func(operation.PatchOperation) error) error {
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		_ = ctx
		if _, ok := seen[value.ID]; ok { continue }
		seen[value.ID] = struct{}{}
		if err := operation.Validate(value); err != nil { return err }
		if err := deliver(value.Clone()); err != nil { return err }
	}
	return nil
}
