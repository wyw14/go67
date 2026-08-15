package patcher

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"
	"example.com/go67/operation"
	"example.com/go67/document"
)

func TestProcessCancellationAndDeduplication(t *testing.T) {
	ctx,cancel:=context.WithCancel(context.Background()); cancel(); calls:=0
	err:=Process(ctx,[]operation.PatchOperation{{ID:"a",Priority:1}},func(operation.PatchOperation)error{calls++;return nil})
	if !errors.Is(err,context.Canceled)||calls!=0 { t.Fatalf("cancel result err=%v calls=%d",err,calls) }
	calls=0; err=Process(context.Background(),[]operation.PatchOperation{{ID:"a",Priority:1},{ID:"a",Priority:1}},func(operation.PatchOperation)error{calls++;return nil})
	if err!=nil||calls!=1 { t.Fatalf("dedupe err=%v calls=%d",err,calls) }
}

func TestExportTransitionDefaultAndExpiry(t *testing.T) {
	want:=errors.New("encode"); if err:=Export(context.Background(),func()([]byte,error){return nil,want},&bytes.Buffer{}); !errors.Is(err,want){t.Fatalf("lost encoder error: %v",err)}
	v:=operation.PatchOperation{ID:"a",State:"new"}; if err:=Transition(&v,"done"); err==nil||v.State!="new" { t.Fatalf("invalid transition changed state: %#v err=%v",v,err) }
	off:=false; if ResolveEnabled(&off,true) { t.Fatal("explicit false was replaced") }
	s:= document.New(); at:=time.Unix(100,0); s.Save(operation.PatchOperation{ID:"a",UpdatedAt:at}); if Expire(s,at)!=0 { t.Fatal("exact cutoff expired") }
}
