package data

import "testing"

const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}

// go test -bench=. -cpuprofile=cpu.prof
// go tool pprof -http=:8080 cpu.prof
