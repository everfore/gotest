package nil

import (
	"testing"
)

type PersonNil struct {
}

func BenchmarkPersion_nil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = (*PersonNil)(nil)
	}
}
func BenchmarkPersion_new(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = new(PersonNil)
	}
}
func BenchmarkPersion_common(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &PersonNil{}
	}
}

type PersonNil2 struct {
	name   string
	friend [1000]string
}

func BenchmarkPersion_nil_more(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = (*PersonNil2)(nil)
	}
}

func BenchmarkPersion_new_more(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = new(PersonNil2)
	}
}
func BenchmarkPersion_common_more(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &PersonNil2{}
	}
}
