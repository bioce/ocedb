package parser_test

import (
	"oce/parser"
	"testing"
)

var (
	correctSQL = []string{
		"select * from _",
		"select * from a",
		"select * from _t_name",
		"select id from t_name",
		"select id, name from t_name",
		"select * from t_name where id = 3",
		"select * from t_name where age = 3.23",
		"select * from t_name where name = '小明'",
		"select * from t_name where name = \"小明\"",
		"select * from t_name where name = \"小明\" limit 3",
		"insert into mydb (id, name, age) values (3, \"小红\", 18);",
		"select * from mydb WHERE id = 3 and (name = 'nail' or age = 23)",
		"select * from \n mydb where id = 3 and name = ''",
		"insert into dept_emp (Name,sex,age,address,email)values('','','','','');",
	}

	incorrectSQL = []string{
		"select * from t where id = '",
		"select * from t_name where age = 3.2.3",
	}
)

func TestDirect(t *testing.T) {
	for _, sql := range correctSQL {
		_, err := parser.Direct(sql)
		if err != nil {
			t.Errorf("except success but get error %s from sql: %s", err.Error(), sql)
		}
	}
	for _, sql := range incorrectSQL {
		_, err := parser.Direct(sql)
		if err == nil {
			t.Errorf("expect error but success of sql: %s", sql)
		}
	}
}
