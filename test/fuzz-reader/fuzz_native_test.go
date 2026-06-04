package fuzzreader

import "testing"

func FuzzReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		Fuzz(data)
	})
}
