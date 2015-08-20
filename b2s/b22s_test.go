package b2s

import (
	"testing"
)

var (
	v    = "this is string."
	data = []byte(v)
)

func TestBS(t *testing.T) {
	bs := ToBytes(v)
	t.Log(bs)

	str := ToString(bs)
	t.Log(str)
}

func TestBS2(t *testing.T) {
	str := I2String(data)
	t.Log(str)

	bs := I2Bytes(v)
	t.Log(bs)
}

func TestDeep(t *testing.T) {
	bs := ToBytes(v)
	t.Log(bs)
	str := ToString(bs)
	t.Log(str)
	bs2 := ToBytes(str)
	t.Log(bs2)
	str2 := ToString(bs2)
	t.Log(str2)
}

func TestDeep2(t *testing.T) {
	bs := I2Bytes(v)
	t.Log(bs)
	str := I2String(bs)
	t.Log(str)
	bs2 := I2Bytes(str)
	t.Log(bs2)
	str2 := I2String(bs2)
	t.Log(str2)
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToBytes(v)
	}
}
func BenchmarkS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToString(data)
	}
}
func BenchmarkB2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = I2Bytes(v)
	}
}
func BenchmarkS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = I2String(data)
	}
}
func BenchmarkB3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(v)
	}
}
func BenchmarkS3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(data)
	}
}
