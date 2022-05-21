package cosqlparser

import (
	btparser "github.com/blastrain/vitess-sqlparser/sqlparser"
	pcparser "github.com/pingcap/parser"
	_ "github.com/pingcap/parser/test_driver"
	vtparser "github.com/wind-c/cosqlparser/sqlparser"
	"testing"
)

var sql = "SELECT cola a, colb b FROM topic t  WHERE b > 20 and a between 10 and 20 GROUP BY (a, b) HAVING a > c ORDER BY b"

//var sql = "SELECT cola, colb FROM topic t  WHERE colb > 20"

func BenchmarkBlastrainVitessParse(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		btparser.Parse(sql)
	}
}

func BenchmarkVitessParse(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vtparser.Parse(sql)
	}
}

func BenchmarkPingcapParse(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pcparser.New().Parse(sql, "", "")
	}
}
