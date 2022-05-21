package sqlparser

import "testing"

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("SELECT cola a, colb b FROM topic t  WHERE b > 20 and a between 10 and 20 GROUP BY (a, b) HAVING a > c ORDER BY b")
	}
}
