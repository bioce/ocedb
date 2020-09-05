%{

package sqlparser

import (
	"oce/parser/ast"
)

%}

%union {
    TableName string
    FieldName string
    str string
}

%type <str> string
%token 'select' 'from' 'where'
%token TABLE_NAME FIELD_NAME
%token <str> STRING

%%
top :       'select' string 'from' string 'where'
                {
                    if l, ok := yylex.(*simpleLex); ok {
                        l.Stmt = &ast.SelectStmt{
                            TableName: $2,
                            FieldName: $4,
                        }
                    }
                }
string :    STRING
                { $$ = $1 }
           ;
%%
