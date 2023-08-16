package test_62049

import "testing"

type Option interface{ option() }

type flags uint64
func (flags) option() {}

func WithRed(v bool) Option {
	return flags(btoi(v) << 15)
}

func btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func WithGreen(v bool) Option {
	if v {
		return flags(1 << 15)
	} else {
		return flags(0)
	}
}

var sink Option

func BenchmarkRed(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = WithRed(true)
	}
}

func BenchmarkGreen(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = WithGreen(true)
	}
}