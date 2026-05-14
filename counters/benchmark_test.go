package counters

import (
	"testing"
	"time"
)

func BenchmarkCounterByDays(b *testing.B) {
	now := time.Now()
	obj := NewCounterByDays(7, now)
	for i := 0; i < b.N; i++ {
		obj.Add(int32(i), now)
	}
}
