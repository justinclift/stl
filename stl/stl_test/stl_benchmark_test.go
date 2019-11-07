package stl_test

import (
	"runtime"
	"strconv"
	"testing"

	"gitlab.com/justinclift/stl/stl"
)

func BenchmarkFrom(b *testing.B) {
	for _, testLevel := range []int{
		// Threads allowed for execution. This will then get used by concurrencyLevel to determine how many worker goroutines are made
		// The number of cores (x2 for hyper-threading) seem to get the best performance
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 32, 40, 48, 56, 64,
	} {
		b.Run("cl="+strconv.Itoa(testLevel), func(b *testing.B) {
			runtime.GOMAXPROCS(testLevel)
			for i := 0; i < b.N; i++ {
				// Read into blank identifier as the actual output does not matter
				_, err := stl.FromFile("testdata/Utah_teapot.stl")
				if err != nil {
					b.Errorf("could not read stl: %v", err)
				}
			}
		})
	}
}
