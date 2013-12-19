package b

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"testing"
)

// compare function
func cmp(a, b []byte) int {
	return bytes.Compare(a, b)
}

// set bench
func BenchmarkSet(b *testing.B) {
	t := TreeNew(cmp)
	debug.FreeOSMemory()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Set([]byte(fmt.Sprintf("k%d", i)), []byte(fmt.Sprintf("v%d", i)))
	}
}

// get bench
func BenchmarkGet(b *testing.B) {
	t := TreeNew(cmp)
	for i := 0; i < b.N; i++ {
		t.Set([]byte(fmt.Sprintf("k%d", i)), []byte(fmt.Sprintf("v%d", i)))
	}
	debug.FreeOSMemory()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Get([]byte(fmt.Sprintf("k%d", i)))
	}
}

// del bench
func BenchmarkDel(b *testing.B) {
	t := TreeNew(cmp)
	for i := 0; i < b.N; i++ {
		t.Set([]byte(fmt.Sprintf("k%d", i)), []byte(fmt.Sprintf("v%d", i)))
	}
	debug.FreeOSMemory()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Delete([]byte(fmt.Sprintf("k%d", i)))
	}
}

// seek bench
func BenchmarkSeek(b *testing.B) {
	t := TreeNew(cmp)
	for i := 0; i < b.N; i++ {
		t.Set([]byte(fmt.Sprintf("k%d", i)), []byte(fmt.Sprintf("v%d", i)))
	}
	debug.FreeOSMemory()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Seek([]byte(fmt.Sprintf("k%d", i)))
	}
}
