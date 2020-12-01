package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {

	t.Run("Should fetch correct result part 1", func(t *testing.T) {
		got := GetResult(2)
		assert.Equal(t, 744475, got)
	})

	t.Run("Should fetch correct result part 2", func(t *testing.T) {
		got := GetResult(3)
		assert.Equal(t, 70276940, got)
	})

}

func BenchmarkTwoMatches(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetResult(2)
	}
}

// BenchmarkTwoMatches-4   	   46119	     24370 ns/op	    9120 B/op	     214 allocs/op
// ~ 0.025ms

func BenchmarkThreeMatches(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetResult(3)
	}
}

// V1
// BenchmarkThreeMatches-4   	    1364	    903440 ns/op	    9120 B/op	     214 allocs/op
// ~ 1ms

// V2
// BenchmarkThreeMatches-4   	    3306	    406025 ns/op	    9120 B/op	     214 allocs/op
// ~0.5 ms
