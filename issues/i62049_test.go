/*
https://go.dev/play/p/bdfBklZm5TE?v=gotip
*/
package test_62049


import "testing"

type Option interface{ option() }
type flags uint64

func (flags) option() {}

func WithRed(v bool) Option {
	return flags(btoi(v))
}
func btoi(b bool) int {
	if b {
		return 1 << 15
	} 
	return 2
}

func WithGreen(v bool) Option {
	if v {
		return flags(1<<15)
	} 
	return flags(2)
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