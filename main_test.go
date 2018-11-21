package main

import "testing"

func TestLoop(t *testing.T) {
	var tests = []struct {
		s        string
		expected int
	}{
		{"----+-+-+-+----++-+-++++-++-+---+--+++++-+++++-++-++++-+++---++-+--+----++-+++---+--+----+++---+-+-+", 51},
		{"--+-", 3},
		{"+++", 0},
	}
	for _, test := range tests {
		if result := loop(test.s); result != test.expected {
			t.Error("Test Failed: Expected {} to equal {}, recieved: {}", test.s, test.expected, result)
		} else {
			t.Logf("Test passed")
		}
	}
}

func benchmarkLoop(s string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		loop(s)
	}
}

func BenchmarkLoop1(b *testing.B) {
	benchmarkLoop("----+-+-+-+----++-+-++++-++-+---+--+++++-+++++-++-++++-+++---++-+--+----++-+++---+--+----+++---+-+-+", b)
}
func BenchmarkLoop2(b *testing.B) { benchmarkLoop("--+-", b) }
func BenchmarkLoop3(b *testing.B) { benchmarkLoop("+++", b) }
