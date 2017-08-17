package itoa

import (
	"fmt"
	"strconv"
	"testing"
	// "time"
)

func FmtItoa(i int64) string {
	return fmt.Sprintf("%d", i)
}

func ConvItoa(i int64) string {
	return strconv.FormatInt(i, 10)
}

func TestItoa(t *testing.T) {
	i := int64(100001)
	fi := FmtItoa(i)
	ci := ConvItoa(i)
	t.Logf("%d, %s, %s", i, fi, ci)
}

func BenchmarkFmtItoa(b *testing.B) {
	it := int64(100001)
	for i := 0; i < b.N; i++ {
		_ = FmtItoa(it)
	}
}

func BenchmarkConvItoa(b *testing.B) {
	it := int64(100001)
	for i := 0; i < b.N; i++ {
		_ = ConvItoa(it)
	}
}

// BenchmarkFmtItoa-4    	10000000	       144 ns/op
// BenchmarkConvItoa-4   	30000000	        48.2 ns/op
