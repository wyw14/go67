package patcher

import (
	"time"
	"example.com/go67/document"
)

func Expire(store *document.Store, cutoff time.Time) int { return store.DeleteBefore(cutoff) }
