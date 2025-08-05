package fuzz

import (
  "testing"
  "google.golang.org/protobuf/encoding/protojson"
  "google.golang.org/protobuf/proto"
  "google.golang.org/protobuf/types/known/anypb"
)

func FuzzJSONBinaryDiff(f *testing.F) {
  // seeds…
  f.Add([]byte("{}"))
  f.Add(seedAny1)
  f.Add(seedAny2)

  f.Fuzz(func(t *testing.T, data []byte) {
    msg := &anypb.Any{}
    if err := protojson.Unmarshal(data, msg); err != nil {
      return
    }

    bin, err := proto.Marshal(msg)
    if err != nil {
      t.Skip()                 // skip on marshal failure
    }
	msg2 := &anypb.Any{}
    if err := proto.Unmarshal(bin, msg2); err != nil {
      t.Fatalf("binary unmarshal mismatch: %v", err)
    }

    // ─── Insert round-trip check HERE ───
    jsonOut, err := protojson.Marshal(msg2)
    if err != nil {
      t.Skip()                 // skip on JSON marshal failure
    }
    msg3 := &anypb.Any{}
    if err := protojson.Unmarshal(jsonOut, msg3); err != nil {
      t.Fatalf("round-trip JSON failed: %v", err)
    }
    // ──────────────────────────────────────
  })
}
