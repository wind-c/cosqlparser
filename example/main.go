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
