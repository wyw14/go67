package operation

import "errors"

var ErrInvalid = errors.New("invalid patcher value")

func Validate(e PatchOperation) error {
	if e.ID == "" { return ErrInvalid }
	if e.Priority < 0 { return ErrInvalid }
	return nil
}
