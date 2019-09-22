package main

import "testing"

func BenchmarkNumFiles(b *testing.B) {
	b.N = 43
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		main()
	}

}
