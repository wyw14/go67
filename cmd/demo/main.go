package main

import (
	"context"
	"fmt"
	"time"
	"example.com/go67/operation"
	"example.com/go67/patcher"
)

func main() {
	v := operation.PatchOperation{ID:"demo", State:"new", UpdatedAt:time.Now()}
	err := patcher.Process(context.Background(), []operation.PatchOperation{v}, func(got operation.PatchOperation) error { fmt.Println(got.ID); return nil })
	if err != nil { panic(err) }
}
