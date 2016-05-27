package capable

import (
	"strings"
	"testing"
)

var s = strings.Repeat("a", 1024)

func standardStr2byte() {
	b := []byte(s)
	_ = string(b)
}

func unsafeStr2Byte() {
	b := str2byte(s)
	_ = byte2str(b)
}

func Test_str2byte(t *testing.T) {
	b := []byte(s)
	B := str2byte(s)
	if len(b) != len(B) {
		t.FailNow()
	}
	size := len(b)
	for i := 0; i < size; i++ {
		if b[i] != B[i] {
			t.FailNow()
		}
	}
}

func Test_byte2str(t *testing.T) {
	b := []byte(s)
	s := string(b)

	B := str2byte(s)
	S := byte2str(B)
	if len(s) != len(S) {
		t.FailNow()
	}
	size := len(b)
	for i := 0; i < size; i++ {
		if s[i] != S[i] {
			t.FailNow()
		}
	}
}

func BenchmarkTest_str2byte_standard(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		standardStr2byte()
	}
}

func BenchmarkTest_str2byte_unsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		unsafeStr2Byte()
	}
}
