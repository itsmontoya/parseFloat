package parseFloat

import (
	"strconv"
	"testing"
)

var val float64
var testVal = []byte("123456")

func TestBasic(t *testing.T) {
	Parse(testVal)
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val = Parse(testVal)
	}

	b.ReportAllocs()
}

func BenchmarkStdlibInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ = strconv.ParseFloat(string(testVal), 64)
	}

	b.ReportAllocs()
}
