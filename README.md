# cosqlparser

Cosqlparser is a powerful and high-performance SQL parser for Go (powered by vitess).

This library inspired by https://github.com/blastrain/vitess-sqlparser, but it uses an older version of vitess and the latest version has vastly improved performance.

(original source : https://github.com/youtube/vitess/tree/master/go/vt/sqlparser)

# Benchmark

### Simple SQL
sql = "SELECT cola, colb FROM topic t  WHERE colb > 20"

|      library       |   b.N    |    ns/op    |    B/op    |  allocs/op   |
|:------------------:|:--------:|:-----------:|:----------:|:------------:|
| wind-c/cosqlparser |  243424  | 4890 ns/op  | 1208 B/op  | 22 allocs/op |
|  vitess-sqlparser  |  158748  | 7431 ns/op  | 11984 B/op | 38 allocs/op |
|   pingcap/parser   |  202621  | 6317 ns/op  | 19064 B/op | 28 allocs/op |

### Complex SQL
sql = "SELECT cola a, colb b FROM topic t  WHERE b > 20 and a between 10 and 20 GROUP BY (a, b) HAVING a > c ORDER BY b"

|      library       |   b.N   |    ns/op    |    B/op    |  allocs/op   |
|:------------------:|:-------:|:-----------:|:----------:|:------------:|
| wind-c/cosqlparser | 107509  | 11048 ns/op | 2305 B/op  | 47 allocs/op |
|  vitess-sqlparser  | 83821   | 14271 ns/op | 13120 B/op | 89 allocs/op |
|   pingcap/parser   | 113925  | 10483 ns/op | 21888 B/op | 60 allocs/op |

# Installation

## [NOTE] Required Go version more than 1.9

```
go get -u github.com/wind-c/cosqlparser
```

# Examples

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/wind-c/cosqlparser/sqlparser"
)

func main() {
	sqlStr := "SELECT t.cola a, t.colb b FROM topic t  WHERE a < 30 and b > 20 and a between 10 and 20 GROUP BY (a, b) HAVING a > c ORDER BY b"

	stmt, err := sqlparser.Parse(sqlStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stmt = %+v\n", stmt)

	jn, err := json.Marshal(stmt)
	if err != nil {
		fmt.Printf("unexpected error: %s", err)
		return
	}
	fmt.Println(string(jn))
}

```