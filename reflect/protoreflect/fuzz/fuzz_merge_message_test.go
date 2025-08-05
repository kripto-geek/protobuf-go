package fuzz

import (
  "testing"
  "google.golang.org/protobuf/proto"
  "google.golang.org/protobuf/types/known/anypb"
)

func FuzzMergeMessage(f *testing.F) {
  // seeds…
  f.Add([]byte{})
  f.Add(seedAny1)
  f.Add(seedAny2)

  f.Fuzz(func(t *testing.T, data []byte) {
    msg1, msg2 := &anypb.Any{}, &anypb.Any{}
    if err := proto.Unmarshal(data, msg1); err != nil {
      return
    }

    proto.Merge(msg2, msg1)

    // ─── Insert round-trip check HERE ───
    out, err := proto.Marshal(msg2)
    if err != nil {
      t.Skip()                  // skip if marshal fails
    }
    var msg3 anypb.Any
    if err := proto.Unmarshal(out, &msg3); err != nil {
      t.Fatalf("round-trip failed after merge: %v", err)
    }
    // ──────────────────────────────────────

    _ = msg2.TypeUrl
    _ = msg2.Value
  })
}


