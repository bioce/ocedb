package parser

import (
	"fmt"
	"log"
	"oce/parser/ast"
)

type Parser struct {
	lex Lexer
}

//func NewParser(sql string) *Parser {
//	return &Parser{lex:NewLexer(sql)}
//}

func Parse(sql string) (stmt ast.Stmt, err error) {
	lex, e := NewLexer(sql)
	if e != nil {
		log.Fatalf("NewLexer(%s) error: %e", sql, e)
		return stmt, e
	}
	tok, _, err := lex.Next()
	if err != nil {
		return nil, err
	}
	if tok == SELECT {
		return parseSelectStmp(lex)
	}
	return nil, fmt.Errorf("parse error sql: %s", sql)
}

func parseSelectStmp(lex Lexer) (ast.Stmt, error) {
	stmt := ast.SelectStmt{}
	fields, e := parseFields(lex)
	if e != nil {
		return nil, e
	}
	stmt.Fields = fields
	return stmt, nil
}

func parseFields(lex Lexer) (ast.Fields, error) {
	tok, str, err := lex.Next()
	if err != nil {
		return nil, err
	}
	if FUZZY == tok {
		return ast.Fields{{AllFields:true}}, nil
	}
	fields := ast.Fields{}
	for {
		fields = append(fields, ast.Field{Name:str})
		tok, str, err = lex.Next()
		if err != nil {
			return nil, err
		}
		if COMMA == tok {
			continue
		} else if FROM == tok {
			lex.Back()
			return fields, nil
		}
		return nil, fmt.Errorf("parseFields expect , or select but get: %s", str)
	}
}

func Expect(lex Lexer, token Token) error {
	tok, _, err := lex.Next()
	if err != nil {
		return err
	}
	if token != tok {
		return fmt.Errorf("Expect token %s get %s ", tokenMapping[token], tokenMapping[tok])
	}
	return nil
}
