package mathutil

import (
	"fmt"
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	fmt.Println(Round(1.23456, 1))
	fmt.Println(Round(1.23456, 2))
	fmt.Println(Round(1.23456, 3))
	fmt.Println(Round(1.23456, 4))
	fmt.Println(Round(1.44445, 4))
}

func Benchmark_Pow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(10, 10)
	}
}

func Benchmark_PowInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowInt(10, 10)
	}
}

//182->31 because intpow func
func Benchmark_Round1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Round(1.23456789, 10)
	}
}
