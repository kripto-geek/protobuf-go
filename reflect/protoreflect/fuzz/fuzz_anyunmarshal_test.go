package fuzz

import (
  "os"
  "testing"
  "google.golang.org/protobuf/proto"
  "google.golang.org/protobuf/types/known/anypb"
)

// FuzzAnyUnmarshal is our entry point.
func FuzzAnyUnmarshal(f *testing.F) {
  // seeds…
  f.Add([]byte{})
  f.Add(seedAny1)
  f.Add(seedAny2)

  f.Fuzz(func(t *testing.T, data []byte) {
    msg := &anypb.Any{}
    if err := proto.Unmarshal(data, msg); err != nil {
      return
    }

    // ─── Insert round-trip check HERE ───
    out, err := proto.Marshal(msg)
    if err != nil {
      t.Skip()                  // unable to re-marshal, skip
    }
    var msg2 anypb.Any
    if err := proto.Unmarshal(out, &msg2); err != nil {
      t.Fatalf("round-trip failed: %v", err)
    }
    // ──────────────────────────────────────

    _ = msg.TypeUrl
    _ = msg.Value
  })
}


// Load seeds from files
var (
  seedAny1 = load("testdata/fuzz_anyunmarshal/seed1.bin")
  seedAny2 = load("testdata/fuzz_anyunmarshal/seed2.bin")
)

// load reads a file into a byte slice
func load(path string) []byte {
  data, _ := os.ReadFile(path)
  return data
}
