package patcher

import (
	"context"
	"io"
)

func Export(ctx context.Context, encode func() ([]byte,error), dst io.Writer) error {
	select { case <-ctx.Done(): return ctx.Err(); default: }
	payload, err := encode()
	if err != nil { return err }
	_, err = dst.Write(payload)
	return err
}
