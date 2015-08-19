package nil

import (
	"testing"
)

type PersonNil struct {
}

func BenchmarkPersionNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = new(PersonNil)
	}
}
func BenchmarkPersionCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &PersonNil{}
	}
}

func BenchmarkPersionNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = (*PersonNil)(nil)
	}
}

type PersonNil2 struct {
	name   string
	friend [101]string
}

func BenchmarkPersionNew2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = new(PersonNil2)
	}
}
func BenchmarkPersionCommon2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &PersonNil2{}
	}
}

func BenchmarkPersionNil2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = (*PersonNil2)(nil)
	}
}
