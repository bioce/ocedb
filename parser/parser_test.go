package parser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var sql string
	sql = "select * from t_name"
	stmt, err := Parse(sql)
	if err != nil {
		t.Errorf("sql: %s e: %e", sql, err)
	} else {
		fmt.Printf("sql: %s stmt: %v\n", sql, stmt)
	}
	//Parse("select * from t_name where id = 3")
	//Parse("select id, name from t_name where id = 3")
}
