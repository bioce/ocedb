package parser

import "strings"

type Token int

const (
	INVALID Token = iota
	UNKNOWN
	EOF

	identifierStart
	LT
	LE
	EQ
	NE
	GT
	GE
	FUZZY
	FROM
	WHERE
	COMMA
	SELECT
	identifierEnd
)

var (
	identifierMaps map[string]Token
	tokenMapping = map[Token]string{
		LT: "<",
		LE: "<=",
		EQ: "=",
		NE: "!=",
		GT: ">",
		GE: ">=",
		FUZZY: "*",
		FROM: "from",
		WHERE: "where",
		COMMA: ",",
		SELECT: "select",
	}
)

func init() {
	identifierMaps = make(map[string]Token)
	for i := identifierStart+1; i < identifierEnd; i++ {
		identifierMaps[tokenMapping[i]] = i
		identifierMaps[strings.ToUpper(tokenMapping[i])] = i
	}
}

func GetToken(lit string) Token {
	if token, ok := identifierMaps[lit]; ok {
		return token
	}
	return UNKNOWN
}
