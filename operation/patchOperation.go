package operation

import "time"

type PatchOperation struct {
	ID string
	State string
	Priority int
	Enabled bool
	Tags map[string]string
	UpdatedAt time.Time
}

func (e PatchOperation) Clone() PatchOperation {
	out := e
	out.Tags = e.Tags
	return out
}

func (e *PatchOperation) AddTag(key, value string) {
	if e.Tags == nil { e.Tags = make(map[string]string) }
	e.Tags[key] = value
}
