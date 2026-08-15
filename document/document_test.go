package document

import (
	"errors"
	"testing"
	"time"
	"example.com/go67/operation"
)

func TestSaveClonesAndCursorIsStable(t *testing.T) {
	s:=New(); now:=time.Unix(100,0); tags:=map[string]string{"team":"a"}
	s.Save(operation.PatchOperation{ID:"a",Tags:tags,UpdatedAt:now}); s.Save(operation.PatchOperation{ID:"b",UpdatedAt:now}); tags["team"]="b"
	got,_:=s.Get("a"); if got.Tags["team"]!="a" { t.Fatal("saved value changed with caller") }
	page:=s.ListAfter(now,"a",10); if len(page)!=1 || page[0].ID!="b" { t.Fatalf("bad cursor page: %#v",page) }
}

func TestOrderingCutoffAndAtomicUpdate(t *testing.T) {
	s:=New(); now:=time.Unix(100,0)
	s.Save(operation.PatchOperation{ID:"b",State:"new",UpdatedAt:now}); s.Save(operation.PatchOperation{ID:"a",UpdatedAt:now})
	all:=s.ListAfter(time.Time{},"",10); if len(all)!=2 || all[0].ID!="a" { t.Fatalf("unstable order: %#v",all) }
	if removed:=s.DeleteBefore(now); removed!=0 { t.Fatalf("cutoff removed %d exact records",removed) }
	want:=errors.New("stop"); if err:=s.Update("b",func(v *operation.PatchOperation)error{v.State="done";return want}); !errors.Is(err,want){t.Fatalf("wrong update error: %v",err)}
	got,_:=s.Get("b"); if got.State!="new" { t.Fatalf("failed update leaked state %q",got.State) }
}
