package ast

type Stmt interface {

}

type Field struct {
	Table Table
	AllFields bool
	Alias string
	Name string
}

type Table struct {
	Alias string
	Name string
}

type Condition struct {

}

type Fields []Field

type Tables []Table

type Conditions []Condition

type SelectStmt struct {
	Tables Tables
	Fields Fields
	Conditions Conditions
}

type InsertStmt struct {

}

type CreateStmt struct {

}
