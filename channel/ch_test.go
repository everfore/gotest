package channel

import (
	"testing"
)

var (
	chb   chan bool
	chst  chan struct{}
	chbig chan [1 << 15]byte
	bigbs [1 << 15]byte
)

func init() {
	chb = make(chan bool, 100)
	chst = make(chan struct{}, 100)
	chbig = make(chan [1 << 15]byte, 100)
}

func BenchmarkChb2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chb <- true
		_ = <-chb
	}
}

func BenchmarkChst2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chst <- struct{}{}
		_ = <-chst
	}
}
func BenchmarkChbig2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chbig <- bigbs
		_ = <-chbig
	}
}

func BenchmarkChb(b *testing.B) {
	b.StopTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			_ = <-chb
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		chb <- true
	}
}

func BenchmarkChst(b *testing.B) {
	b.StopTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			_ = <-chst
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		chst <- struct{}{}
	}
}

func BenchmarkChbig(b *testing.B) {
	b.StopTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			_ = <-chbig
		}
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		chbig <- bigbs
	}
}
